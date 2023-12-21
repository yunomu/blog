package filedb

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const (
	attr_ORIGIN         = "orig"
	attr_REPLICA_FORMAT = "rep:%s"
)

type DynamoDBRecord struct {
	Key         string `dynamodbav:"Key"`
	Attr        string `dynamodbav:"Attr"`
	UserId      string `dynamodbav:"UserId,omitempty"`
	ContentType string `dynamodbav:"CType,omitempty"`
	Name        string `dynamodbav:"Name,omitempty"`
	Timestamp   int64  `dynamodbav:"TS,omitempty"`
	Size        int    `dynamodbav:"Size,omitempty"`
	Width       int    `dynamodbav:"W,omitempty"`
	Height      int    `dynamodbav:"H,omitempty"`
	Status      string `dynamodbav:"Status,omitmpty"`
	UserIdx     string `dynamodbav:"UserIdx,omitempty"`
}

type DynamoDBLogger interface {
	Error(msg string)
}

type defaultLogger struct{}

func (l *defaultLogger) Error(msg string) {}

type DynamoDB struct {
	client    dynamodbiface.DynamoDBAPI
	table     string
	userIndex string

	logger DynamoDBLogger
}

var _ DB = (*DynamoDB)(nil)

type DynamoDBOption func(*DynamoDB)

func SetDynamoDBLogger(l DynamoDBLogger) DynamoDBOption {
	return func(db *DynamoDB) {
		if l != nil {
			db.logger = l
		} else {
			db.logger = &defaultLogger{}
		}
	}
}

func NewDynamoDB(
	client dynamodbiface.DynamoDBAPI,
	table string,
	userIndex string,
	options ...DynamoDBOption,
) *DynamoDB {
	db := &DynamoDB{
		client:    client,
		table:     table,
		userIndex: userIndex,

		logger: &defaultLogger{},
	}
	for _, f := range options {
		f(db)
	}

	return db
}

func getTimestamp() int64 {
	return time.Now().UnixMilli()
}

func (db *DynamoDB) Reserve(ctx context.Context, key, userId string) (int64, error) {
	ts := getTimestamp()
	item, err := dynamodbattribute.MarshalMap(&DynamoDBRecord{
		Key:       key,
		Attr:      attr_ORIGIN,
		Name:      attr_ORIGIN,
		UserId:    userId,
		UserIdx:   userId,
		Status:    Status_RESERVED,
		Timestamp: ts,
	})
	if err != nil {
		db.logger.Error("marshal error at create")
		return 0, err
	}

	expr, err := expression.NewBuilder().
		WithCondition(expression.AttributeNotExists(expression.Name("Attr"))).
		Build()
	if err != nil {
		db.logger.Error("create expression error at create")
		return 0, err
	}

	if _, err := db.client.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.table),
		Item:      item,

		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}); err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				return 0, ErrConflictKey
			}
		}
		db.logger.Error("put item error at create")
		return 0, err
	}

	return ts, nil
}

func (db *DynamoDB) CreateCommit(ctx context.Context, key, contentType string, timestamp int64, size, width, height int) error {
	expr, err := expression.NewBuilder().
		WithCondition(expression.Equal(expression.Name("TS"), expression.Value(timestamp))).
		WithUpdate(
			expression.Set(expression.Name("CType"), expression.Value(contentType)).
				Set(expression.Name("Size"), expression.Value(size)).
				Set(expression.Name("W"), expression.Value(width)).
				Set(expression.Name("H"), expression.Value(height)).
				Set(expression.Name("Status"), expression.Value(Status_AVAILABLE)),
		).
		Build()
	if err != nil {
		db.logger.Error("create expression error at create")
		return err
	}

	if _, err := db.client.UpdateItemWithContext(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(db.table),
		Key: map[string]*dynamodb.AttributeValue{
			"Key": {S: aws.String(key)},
		},

		UpdateExpression:          expr.Update(),
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}); err != nil {
		return err
	}

	return nil
}

func (db *DynamoDB) CreateReplica(ctx context.Context, key, name, userId, contentType string, size, width, height int) error {
	ts := getTimestamp()
	item, err := dynamodbattribute.MarshalMap(&DynamoDBRecord{
		Key:       key,
		Attr:      fmt.Sprintf(attr_REPLICA_FORMAT, name),
		Name:      name,
		UserId:    userId,
		Status:    Status_AVAILABLE,
		Timestamp: ts,
		Size:      size,
		Width:     width,
		Height:    height,
	})
	if err != nil {
		db.logger.Error("marshal error at create replica")
		return err
	}

	expr, err := expression.NewBuilder().
		WithCondition(expression.AttributeNotExists(expression.Name("Attr"))).
		Build()
	if err != nil {
		db.logger.Error("build expression error at create replica")
		return err
	}

	if _, err := db.client.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.table),
		Item:      item,

		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}); err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				return ErrConflictKey
			}
		}
		db.logger.Error("put item error at create replica")
		return err
	}

	return nil
}

func (db *DynamoDB) List(ctx context.Context, userId string, options ...ListOption) ([]*File, string, error) {
	opts := &listOptions{
		limit: 10,
	}
	for _, f := range options {
		f(opts)
	}

	expr, err := expression.NewBuilder().
		WithKeyCondition(
			expression.KeyEqual(expression.Key("UserIdx"), expression.Value(userId)),
		).
		Build()
	if err != nil {
		db.logger.Error("create expression error at list")
		return nil, "", err
	}

	var startKey map[string]*dynamodb.AttributeValue
	if opts.continuationToken != "" {
		ts, err := strconv.ParseInt(opts.continuationToken, 10, 64)
		if err != nil {
			return nil, "", ErrInvalidContinuationToken
		}

		key, err := dynamodbattribute.MarshalMap(DynamoDBRecord{
			UserId:    userId,
			Timestamp: ts,
		})
		if err != nil {
			return nil, "", err
		}

		startKey = key
	}

	var lastKey string
	var ret []*File
	var rerr error
	if err := db.client.QueryPagesWithContext(ctx, &dynamodb.QueryInput{
		TableName: aws.String(db.table),
		IndexName: aws.String(db.userIndex),

		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),

		Limit:             aws.Int64(int64(opts.limit)),
		ScanIndexForward:  aws.Bool(false),
		ExclusiveStartKey: startKey,
	}, func(out *dynamodb.QueryOutput, last bool) bool {
		var recs []DynamoDBRecord
		if err := dynamodbattribute.UnmarshalListOfMaps(out.Items, &recs); err != nil {
			rerr = err
			return false
		}

		for _, r := range recs {
			if r.Status != Status_AVAILABLE {
				continue
			}

			ret = append(ret, &File{
				Key:         r.Key,
				UserId:      r.UserId,
				ContentType: r.ContentType,
				Entities: []*Entity{
					{
						Name:      r.Name,
						Timestamp: r.Timestamp,
						Size:      r.Size,
						Width:     r.Width,
						Height:    r.Height,
						Status:    r.Status,
					},
				},
			})
		}

		if last {
			var key DynamoDBRecord
			if err := dynamodbattribute.UnmarshalMap(out.LastEvaluatedKey, &key); err != nil {
				rerr = err
				return false
			}

			lastKey = strconv.FormatInt(key.Timestamp, 10)
		}

		return true
	}); rerr != nil {
		db.logger.Error("query inner error at list")
		return nil, "", rerr
	} else if err != nil {
		db.logger.Error("query error at list")
		return nil, "", err
	}

	return ret, lastKey, nil
}

func (db *DynamoDB) getItems(ctx context.Context, key string) ([]DynamoDBRecord, error) {
	expr, err := expression.NewBuilder().
		WithKeyCondition(
			expression.KeyEqual(expression.Key("Key"), expression.Value(key)),
		).
		Build()
	if err != nil {
		db.logger.Error("create expression error at get")
		return nil, err
	}

	var ret []DynamoDBRecord
	var rerr error
	if err := db.client.QueryPagesWithContext(ctx, &dynamodb.QueryInput{
		TableName: aws.String(db.table),

		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}, func(out *dynamodb.QueryOutput, last bool) bool {
		var recs []DynamoDBRecord
		if err := dynamodbattribute.UnmarshalListOfMaps(out.Items, &recs); err != nil {
			rerr = err
			return false
		}

		ret = append(ret, recs...)

		return true
	}); rerr != nil {
		db.logger.Error("query inner error at get")
		return nil, rerr
	} else if err != nil {
		db.logger.Error("query error at get")
		return nil, err
	}

	return ret, nil
}

func (db *DynamoDB) Get(ctx context.Context, key string) (*File, error) {
	recs, err := db.getItems(ctx, key)
	if err != nil {
		return nil, err
	}

	var ret File
	for _, r := range recs {
		if r.Status != Status_AVAILABLE {
			continue
		}

		ret.Key = r.Key
		ret.UserId = r.UserId
		ret.ContentType = r.ContentType
		ret.Entities = append(ret.Entities, &Entity{
			Name:      r.Name,
			Timestamp: r.Timestamp,
			Size:      r.Size,
			Width:     r.Width,
			Height:    r.Height,
			Status:    r.Status,
		})
	}
	if len(ret.Entities) == 0 {
		return nil, ErrNotFound
	}

	return &ret, nil
}

func (db *DynamoDB) Delete(ctx context.Context, key, userId string, ts int64) (int64, error) {
	items, err := db.getItems(ctx, key)
	if err != nil {
		return 0, err
	}

	newTS := getTimestamp()
	var writeItems []*dynamodb.TransactWriteItem
	for _, item := range items {
		exprBuilder := expression.NewBuilder().
			WithUpdate(
				expression.Set(expression.Name("Status"), expression.Value(Status_DELETING)).
					Set(expression.Name("TS"), expression.Value(newTS)),
			)
		if item.Attr == attr_ORIGIN {
			exprBuilder = exprBuilder.WithCondition(
				expression.And(
					expression.Equal(expression.Name("UserId"), expression.Value(userId)),
					expression.Equal(expression.Name("TS"), expression.Value(ts)),
				),
			)
		}

		expr, err := exprBuilder.Build()
		if err != nil {
			db.logger.Error("delete expression error at delete")
			return 0, err
		}

		av, err := dynamodbattribute.MarshalMap(DynamoDBRecord{
			Key:  key,
			Attr: item.Attr,
		})
		if err != nil {
			return 0, err
		}

		writeItems = append(writeItems, &dynamodb.TransactWriteItem{
			Update: &dynamodb.Update{
				TableName: aws.String(db.table),
				Key:       av,

				UpdateExpression:          expr.Update(),
				ConditionExpression:       expr.Condition(),
				ExpressionAttributeNames:  expr.Names(),
				ExpressionAttributeValues: expr.Values(),
			},
		})
	}

	if _, err := db.client.TransactWriteItemsWithContext(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: writeItems,
	}); err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				return 0, ErrLock
			}
		}
		db.logger.Error("put items error at delete")
		return 0, err
	}

	return ts, nil
}

func (db *DynamoDB) DeleteCommit(ctx context.Context, key string, ts int64) error {
	items, err := db.getItems(ctx, key)
	if err != nil {
		return err
	}

	exprBuilder := expression.NewBuilder().
		WithCondition(
			expression.Equal(expression.Name("TS"), expression.Value(ts)),
		)

	var writeItems []*dynamodb.TransactWriteItem
	for _, item := range items {
		av, err := dynamodbattribute.MarshalMap(DynamoDBRecord{
			Key:  key,
			Attr: item.Attr,
		})
		if err != nil {
			db.logger.Error("attribute marshal error at delete commit")
			return err
		}

		switch item.Attr {
		case attr_ORIGIN:
			newTS := getTimestamp()
			expr, err := exprBuilder.
				WithUpdate(
					expression.Set(expression.Name("TS"), expression.Value(newTS)).
						Set(expression.Name("Status"), expression.Value(Status_DELETED)).
						Remove(expression.Name("UserId")).
						Remove(expression.Name("Name")).
						Remove(expression.Name("ContentType")).
						Remove(expression.Name("Size")).
						Remove(expression.Name("Width")).
						Remove(expression.Name("Height")).
						Remove(expression.Name("UserIdx")),
				).
				Build()
			if err != nil {
				db.logger.Error("expression error at delete commit")
				return err
			}

			writeItems = append(writeItems, &dynamodb.TransactWriteItem{
				Update: &dynamodb.Update{
					TableName: aws.String(db.table),
					Key:       av,

					UpdateExpression:          expr.Update(),
					ConditionExpression:       expr.Condition(),
					ExpressionAttributeNames:  expr.Names(),
					ExpressionAttributeValues: expr.Values(),
				},
			})
		default:
			expr, err := exprBuilder.Build()
			if err != nil {
				db.logger.Error("expression error at delete commit")
				return err
			}

			writeItems = append(writeItems, &dynamodb.TransactWriteItem{
				Delete: &dynamodb.Delete{
					TableName: aws.String(db.table),
					Key:       av,

					ConditionExpression:       expr.Condition(),
					ExpressionAttributeNames:  expr.Names(),
					ExpressionAttributeValues: expr.Values(),
				},
			})
		}
	}

	return nil
}

package userdb

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type DynamoDBRecord struct {
	Id   string `dynamodbav:"Id"`
	Attr string `dynamodbav:"Attr"`
	Name string `dynamodbav:"Name,omitempty"`
}

type DynamoDBLogger interface {
	Error(msg string)
	NameConflict(*DynamoDBRecord)
}

type defaultLogger struct{}

func (d *defaultLogger) Error(msg string)               {}
func (d *defaultLogger) NameConflict(_ *DynamoDBRecord) {}

type DynamoDB struct {
	client    dynamodbiface.DynamoDBAPI
	table     string
	nameIndex string

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
	nameIndex string,
	options ...DynamoDBOption,
) *DynamoDB {
	ret := &DynamoDB{
		client:    client,
		table:     table,
		nameIndex: nameIndex,
	}
	for _, f := range options {
		f(ret)
	}

	return ret
}

func (db *DynamoDB) checkNameConflict(ctx context.Context, name string) error {
	expr, err := expression.NewBuilder().
		WithKeyCondition(expression.KeyEqual(expression.Key("Name"), expression.Value(name))).
		Build()
	if err != nil {
		return err
	}

	out, err := db.client.QueryWithContext(ctx, &dynamodb.QueryInput{
		TableName: aws.String(db.table),
		IndexName: aws.String(db.nameIndex),

		Limit:  aws.Int64(1),
		Select: aws.String("ALL_PROJECTED_ATTRIBUTES"),

		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		return err
	}

	if len(out.Items) >= 0 {
		var rec DynamoDBRecord
		recs := []DynamoDBRecord{}
		if err := dynamodbattribute.UnmarshalListOfMaps(out.Items, &recs); err != nil {
			db.logger.Error("record unmarshal error at name conflict")
			return err
		} else {
			rec = recs[0]
		}
		db.logger.NameConflict(&rec)

		return ErrNameConflict
	}

	return nil
}

func (db *DynamoDB) Create(ctx context.Context, userId string, name string) (*User, error) {
	if err := db.checkNameConflict(ctx, name); err != nil {
		return nil, err
	}

	item, err := dynamodbattribute.MarshalMap(&DynamoDBRecord{
		Id:   userId,
		Attr: "Main",
		Name: name,
	})
	if err != nil {
		db.logger.Error("record marshal error at create")
		return nil, err
	}

	expr, err := expression.NewBuilder().
		WithCondition(expression.AttributeNotExists(expression.Name("Name"))).
		Build()
	if err != nil {
		db.logger.Error("expression build error at create")
		return nil, err
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
				return nil, ErrAlreadyExist
			}
		}
		db.logger.Error("put item error at create")
		return nil, err
	}

	return &User{
		Id:   userId,
		Name: name,
	}, nil
}

func (db *DynamoDB) Get(ctx context.Context, id string) (*User, error) {
	key, err := dynamodbattribute.MarshalMap(&DynamoDBRecord{
		Id:   id,
		Attr: "Main",
	})
	if err != nil {
		db.logger.Error("record marshal error at get")
		return nil, err
	}

	out, err := db.client.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(db.table),
		Key:       key,
	})
	if err != nil {
		db.logger.Error("get item error at get")
		return nil, err
	}

	if len(out.Item) == 0 {
		return nil, ErrNotFound
	}

	var rec DynamoDBRecord
	if err := dynamodbattribute.UnmarshalMap(out.Item, &rec); err != nil {
		db.logger.Error("record unmarshal error at get")
		return nil, err
	}

	return &User{
		Id:   rec.Id,
		Name: rec.Name,
	}, nil
}

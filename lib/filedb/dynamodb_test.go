package filedb

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestDelete(t *testing.T) {
	db := NewDynamoDB(
		&MockDynamoDB{
			QueryPagesFn: func(in *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool) error {
				fn(&dynamodb.QueryOutput{
					Items: []map[string]*dynamodb.AttributeValue{
						{
							"Key":     {S: aws.String("key")},
							"Attr":    {S: aws.String(attr_ORIGIN)},
							"UserId":  {S: aws.String("user-id")},
							"Name":    {S: aws.String("orig")},
							"CType":   {S: aws.String("image/jpeg")},
							"TS":      {N: aws.String("1")},
							"Size":    {N: aws.String("100")},
							"Status":  {S: aws.String(Status_AVAILABLE)},
							"UserIdx": {S: aws.String("user-id")},
						},
						{
							"Key":    {S: aws.String("key")},
							"Attr":   {S: aws.String("rep:1x2")},
							"UserId": {S: aws.String("user-id")},
							"Name":   {S: aws.String("1x2")},
							"CType":  {S: aws.String("image/jpeg")},
							"TS":     {N: aws.String("1")},
							"Size":   {N: aws.String("100")},
							"Status": {S: aws.String(Status_AVAILABLE)},
						},
					},
				}, true)
				return nil
			},
			TransactWriteItemsFn: func(in *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
				if len(in.TransactItems) != 2 {
					return nil, errors.New("number of items mismatch")
				}

				for _, item := range in.TransactItems {
					update := item.Update
					if update == nil {
						return nil, errors.New("not update request")
					}

					attrAV, ok := update.Key["Attr"]
					if !ok {
						return nil, errors.New("Attr not found")
					}

					if attrAV.S == nil {
						return nil, errors.New("Attr is not string value")
					}
					attr := aws.StringValue(attrAV.S)

					if attr == attr_ORIGIN {
						if update.ConditionExpression == nil {
							return nil, errors.New("condition is empty in origin deletion")
						}
					} else {
						if update.ConditionExpression != nil {
							return nil, errors.New("condition is not empty in replica deletion")
						}
					}
				}

				return &dynamodb.TransactWriteItemsOutput{}, nil
			},
		},
		"table",
		"userIndex",
	)

	ctx := context.Background()
	ts, err := db.Delete(ctx, "key", "user-id", 1)
	if err != nil {
		t.Logf("new timestamp: %v", ts)
		t.Errorf("filedb.Delete: %v", err)
	}
}

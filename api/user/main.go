package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/yunomu/blog/lib/userdb"

	"github.com/yunomu/blog/api/user/internal/handler"
)

var (
	logger *slog.Logger
	debug  bool
)

func init() {
	debug = os.Getenv("DEBUG") != ""

	levelVar := new(slog.LevelVar)
	if debug {
		levelVar.Set(slog.LevelDebug)
	}

	logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: levelVar,
	}))
}

type userdbLogger struct {
	logger *slog.Logger
}

func (l *userdbLogger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *userdbLogger) NameConflict(record *userdb.DynamoDBRecord) {
	l.logger.Error("Name conflict", "conflicted_record", record)
}

type handlerLogger struct {
	logger *slog.Logger
}

func (h *handlerLogger) HandlerError(err error, req *handler.Request) {
	h.logger.Error("Handler", "err", err, "request", req)
}

func (h *handlerLogger) Error(err error, msg string) {
	h.logger.Error(msg, "err", err)
}

func (h *handlerLogger) InitUser(user *userdb.User) {
	h.logger.Info("Initialized user", "user", user)
}

func main() {
	ctx := context.Background()

	table := os.Getenv("USER_TABLE")
	nameIndex := os.Getenv("USER_NAME_INDEX")
	region := os.Getenv("REGION")

	slog.Info("Init",
		"table", table,
		"nameIndex", nameIndex,
		"region", region,
		"debug", debug,
	)

	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		slog.Error("NewSession", "err", err)
		return
	}

	h := handler.NewHandler(
		userdb.NewDynamoDB(
			dynamodb.New(sess),
			table,
			nameIndex,
			userdb.SetDynamoDBLogger(
				&userdbLogger{
					logger: logger.With("module", "userdb"),
				},
			),
		),
		handler.SetLogger(&handlerLogger{
			logger: logger.With("module", "handler"),
		}),
	)

	lambda.StartWithContext(ctx, h.Serve)
}

package main

import (
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/yunomu/blog/lib/cdn"
	"github.com/yunomu/blog/lib/filedb"
	"github.com/yunomu/blog/lib/storage"

	"github.com/yunomu/blog/lambda/image/internal/handler"
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

type filedbLogger struct {
	logger *slog.Logger
}

func (l *filedbLogger) Error(msg string) {
	l.logger.Error(msg)
}

type handlerLogger struct {
	logger *slog.Logger
}

func (h *handlerLogger) ItemFormatError(item interface{}) {
	logger.Error("Item format error", "item", item)
}

func (h *handlerLogger) FileLockError(msg string, file *filedb.File) {
	logger.Info("Filedb lock error", "msg", msg, "file", file)
}

func main() {
	bucket := os.Getenv("BUCKET")
	fileTable := os.Getenv("FILE_TABLE")
	fileUserIndex := os.Getenv("FILE_USER_INDEX")
	distribution := os.Getenv("DISTRIBUTION")
	region := os.Getenv("REGION")

	slog.Info("Init",
		"bucket", bucket,
		"fileTable", fileTable,
		"distribution", distribution,
		"region", region,
		"debug", debug,
	)

	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		slog.Error("NewSession", "err", err)
		return
	}

	h := handler.NewHandler(
		filedb.NewDynamoDB(
			dynamodb.New(sess),
			fileTable,
			fileUserIndex,
			filedb.SetDynamoDBLogger(
				&filedbLogger{logger: logger.With("module", "filedb")},
			),
		),
		storage.NewS3(
			s3.New(sess),
			bucket,
		),
		cdn.NewCloudFront(
			cloudfront.New(sess),
			distribution,
		),
		handler.SetLogger(
			&handlerLogger{logger: logger.With("module", "handler")},
		),
	)

	lambda.Start(h.Serve)
}

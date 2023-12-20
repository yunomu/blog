package cdn

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
)

type CloudFrontLogger interface {
	InvalidationCallerReference(string)
	InvalidationError(err error, distributionId string, callerReference string, paths []string)
	InvalidationCreated(id string, createTime *time.Time, status string)
}

type defaultLogger struct{}

func (d *defaultLogger) InvalidationCallerReference(_ string)              {}
func (d *defaultLogger) InvalidationError(error, string, string, []string) {}
func (d *defaultLogger) InvalidationCreated(string, *time.Time, string)    {}

type CloudFront struct {
	client         cloudfrontiface.CloudFrontAPI
	distributionId string

	logger CloudFrontLogger
}

var _ CDN = (*CloudFront)(nil)

type CloudFrontOption func(*CloudFront)

func SetCloudFrontLogger(l CloudFrontLogger) CloudFrontOption {
	return func(c *CloudFront) {
		if l != nil {
			c.logger = l
		} else {
			c.logger = &defaultLogger{}
		}
	}
}

func NewCloudFront(
	client cloudfrontiface.CloudFrontAPI,
	distributionId string,
	options ...CloudFrontOption,
) *CloudFront {
	ret := &CloudFront{
		client:         client,
		distributionId: distributionId,

		logger: &defaultLogger{},
	}
	for _, f := range options {
		f(ret)
	}

	return ret
}

func (c *CloudFront) Invalidate(ctx context.Context, paths []string) error {
	ref := strconv.FormatInt(time.Now().Unix(), 10)
	c.logger.InvalidationCallerReference(ref)

	out, err := c.client.CreateInvalidationWithContext(ctx, &cloudfront.CreateInvalidationInput{
		DistributionId: aws.String(c.distributionId),
		InvalidationBatch: &cloudfront.InvalidationBatch{
			CallerReference: aws.String(ref),
			Paths: &cloudfront.Paths{
				Items:    aws.StringSlice(paths),
				Quantity: aws.Int64(int64(len(paths))),
			},
		},
	})
	if err != nil {
		c.logger.InvalidationError(err, c.distributionId, ref, paths)
		return err
	}

	c.logger.InvalidationCreated(
		aws.StringValue(out.Invalidation.Id),
		out.Invalidation.CreateTime,
		aws.StringValue(out.Invalidation.Status),
	)

	return nil
}

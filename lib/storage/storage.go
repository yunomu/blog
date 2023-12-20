package storage

import (
	"context"
	"io"
)

type Storage interface {
	Put(ctx context.Context, key string, contentType string, blob io.ReadSeeker) error
}

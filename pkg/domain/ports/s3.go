package ports

import "context"

type S3Port interface {
    ListObjects(ctx context.Context) ([]string, error)
}

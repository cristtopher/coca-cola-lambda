package ports

import "context"

type SecretManagerPort interface {
    GetSecret(ctx context.Context, secretName string) (map[string]string, error)
}

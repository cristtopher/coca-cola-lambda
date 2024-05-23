package ports

import "coca-cola-lambda/internal/core"

type SecretManager interface {
    GetSecret(name string) (*core.Secret, error)
}

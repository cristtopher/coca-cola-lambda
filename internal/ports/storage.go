package ports

type Storage interface {
    GetObject(bucket, key string) (string, error)
}

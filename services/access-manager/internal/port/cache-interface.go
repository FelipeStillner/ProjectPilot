package port

import "time"

type CacheInterface interface {
	Set(key string, value string, time time.Duration) error
	Get(key string) (*string, error)
}

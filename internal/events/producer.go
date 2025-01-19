package events

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CountProducer interface {
	AddCount(ctx context.Context, id int64) error
}

type RedisCountProducer struct {
	cmd redis.Cmdable
	key string
}

// AddCount implements CountProducer.
func (r *RedisCountProducer) AddCount(ctx context.Context, id int64) error {
	return r.cmd.LPush(ctx, r.key, id).Err()
}

func NewRedisCountProducer(cmd redis.Cmdable) CountProducer {
	return &RedisCountProducer{
		cmd: cmd,
		key: "url_counter_queue",
	}
}

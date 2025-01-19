package events

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/chenmuyao/url_shortener/internal/repo/dao"
	"github.com/redis/go-redis/v9"
)

type CountConsumer interface {
	DoAddCount(ctx context.Context, id int64) error
	Start()
}

type RedisCountConsumer struct {
	dao      *dao.Queries
	cmd      redis.Cmdable
	key      string
	maxErr   int
	waitTime time.Duration
}

// DoAddCount implements CountConsumer.
func (r *RedisCountConsumer) DoAddCount(ctx context.Context, id int64) error {
	_, err := r.dao.UpdateCountByID(ctx, id)
	return err
}

func (r *RedisCountConsumer) Start() {
	slog.Info("Consumer started")
	go func() {
		errCnt := 0
		for {
			res, err := r.cmd.BRPop(context.Background(), 0, r.key).Result()
			if err != nil {
				errCnt++
				slog.Error("Redis pop error", slog.Any("err", err))
				if errCnt > r.maxErr {
					break
				}
				time.Sleep(r.waitTime)
			}
			errCnt = 0
			id, err := strconv.ParseInt(res[1], 10, 64)
			if err != nil {
				slog.Error("wrong data", "data", res)
				// should not happen, consider it done...
				continue
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			err = r.DoAddCount(ctx, id)
			if err != nil {
				slog.Error("failed to add count", "id", id, "err", err)
			}
		}
	}()
}

func NewRedisCountConsumer(dao *dao.Queries, cmd redis.Cmdable) CountConsumer {
	return &RedisCountConsumer{
		dao:      dao,
		cmd:      cmd,
		key:      "url_counter_queue",
		maxErr:   10,
		waitTime: 500 * time.Millisecond,
	}
}

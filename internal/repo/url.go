package repo

import (
	"context"
	"log/slog"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/chenmuyao/url_shortener/internal/domain"
	"github.com/chenmuyao/url_shortener/internal/events"
	"github.com/chenmuyao/url_shortener/internal/repo/dao"
	"github.com/jackc/pgx/v5/pgconn"
)

const keyConflictSQLCode = "23505"

type UrlShortenerRepo interface {
	InsertURL(ctx context.Context, full string) (int64, error)
	GetURL(ctx context.Context, id int64) (domain.Url, error)
}

type urlShortenerRepo struct {
	dao     *dao.Queries
	node    *snowflake.Node
	counter events.CountProducer
}

// GetURL implements UrlShortenerRepo.
func (u *urlShortenerRepo) GetURL(ctx context.Context, id int64) (domain.Url, error) {
	res, err := u.dao.GetURLByID(ctx, id)
	go func() {
		c, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		er := u.counter.AddCount(c, id)
		if er != nil {
			// NOTE: must be redis error, log and monitor here
			slog.Error("cannot push add count message", slog.Any("id", id), slog.Any("err", er))
		}
	}()
	return domain.Url(res), err
}

// InsertURL implements UrlShortenerRepo.
func (u *urlShortenerRepo) InsertURL(ctx context.Context, full string) (int64, error) {
	id := u.node.Generate().Int64()

	url, err := u.dao.InsertURL(ctx, dao.InsertURLParams{
		ID:        id,
		Url:       full,
		CreatedAt: time.Now(),
		Count:     0,
	})
	if err != nil {
		return u.handleExisted(ctx, full, err)
	}
	return url.ID, nil
}

func (u *urlShortenerRepo) handleExisted(
	ctx context.Context,
	full string,
	err error,
) (int64, error) {
	pgErr, ok := err.(*pgconn.PgError)
	if ok && pgErr.Code == keyConflictSQLCode {
		id, err := u.dao.GetIDByURL(ctx, full)
		if err != nil {
			return 0, err
		}
		return id, nil
	}

	return 0, err
}

func NewUrlShortenerRepo(
	node *snowflake.Node,
	dao *dao.Queries,
	counter events.CountProducer,
) UrlShortenerRepo {
	return &urlShortenerRepo{dao: dao, node: node, counter: counter}
}

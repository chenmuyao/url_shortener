package repo

import (
	"context"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/chenmuyao/url_shortener/internal/domain"
	"github.com/chenmuyao/url_shortener/internal/repo/dao"
	"github.com/jackc/pgx/v5/pgconn"
)

const keyConflictSQLCode = "23505"

type UrlShortenerRepo interface {
	InsertURL(ctx context.Context, full string) (int64, error)
	GetURL(ctx context.Context, id int64) (domain.Url, error)
}

type urlShortenerRepo struct {
	dao  *dao.Queries
	node *snowflake.Node
}

// GetURL implements UrlShortenerRepo.
func (u *urlShortenerRepo) GetURL(ctx context.Context, id int64) (domain.Url, error) {
	res, err := u.dao.UpdateCountByID(ctx, id)
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

func NewUrlShortenerRepo(dao *dao.Queries, node *snowflake.Node) UrlShortenerRepo {
	return &urlShortenerRepo{dao: dao, node: node}
}

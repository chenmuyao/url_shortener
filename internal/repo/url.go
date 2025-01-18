package repo

import (
	"context"

	"github.com/chenmuyao/url_shortener/internal/repo/dao"
	"github.com/jackc/pgx/v5/pgconn"
)

const keyConflictSQLCode = "23505"

type UrlShortenerRepo interface {
	InsertURL(ctx context.Context, full string) (int64, error)
	GetURL(ctx context.Context, id int64) (string, error)
}

type urlShortenerRepo struct {
	dao *dao.Queries
}

// GetURL implements UrlShortenerRepo.
func (u *urlShortenerRepo) GetURL(ctx context.Context, id int64) (string, error) {
	return u.dao.GetByID(ctx, id)
}

// InsertURL implements UrlShortenerRepo.
func (u *urlShortenerRepo) InsertURL(ctx context.Context, full string) (int64, error) {
	url, err := u.dao.InsertURL(ctx, full)
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

func NewUrlShortenerRepo(dao *dao.Queries) UrlShortenerRepo {
	return &urlShortenerRepo{dao: dao}
}

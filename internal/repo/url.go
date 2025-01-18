package repo

import "context"

type UrlShortenerRepo interface {
	InsertURL(ctx context.Context, full string) (int64, error)
	GetURL(ctx context.Context, id int) (string, error)
}

type urlShortenerRepo struct{}

// GetURL implements UrlShortenerRepo.
func (u *urlShortenerRepo) GetURL(ctx context.Context, id int) (string, error) {
	panic("unimplemented")
}

// InsertURL implements UrlShortenerRepo.
func (u *urlShortenerRepo) InsertURL(ctx context.Context, full string) (int64, error) {
	panic("unimplemented")
}

func NewUrlShortenerRepo() UrlShortenerRepo {
	return &urlShortenerRepo{}
}

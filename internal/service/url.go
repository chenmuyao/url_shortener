package service

import (
	"context"
	"errors"
	"strings"

	"github.com/chenmuyao/url_shortener/internal/domain"
	"github.com/chenmuyao/url_shortener/internal/repo"
)

const base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type UrlShortenerSvc interface {
	// Shorten takes the full url and returns a short url
	Shorten(ctx context.Context, fullUrl string) (string, error)
	// GetCount takes a shortID and returns the domain URL
	GetURL(ctx context.Context, shortID string) (domain.Url, error)
}

type urlShortenerSvc struct {
	repo repo.UrlShortenerRepo
}

// GetFull implements UrlShortenerSvc.
func (u *urlShortenerSvc) GetURL(ctx context.Context, shortID string) (domain.Url, error) {
	id, err := base62Dec(shortID)
	if err != nil {
		return domain.Url{}, err
	}
	return u.repo.GetURL(ctx, id)
}

// Shorten implements UrlShortenerSvc.
func (u *urlShortenerSvc) Shorten(ctx context.Context, fullUrl string) (string, error) {
	id, err := u.repo.InsertURL(ctx, fullUrl)
	if err != nil {
		return "", err
	}

	return base62Enc(id), nil
}

func NewUrlShortenerSvc(repo repo.UrlShortenerRepo) UrlShortenerSvc {
	return &urlShortenerSvc{repo: repo}
}

func base62Enc(num int64) string {
	if num == 0 {
		return string(base62Alphabet[0])
	}
	var encoded string
	for num > 0 {
		remainder := num % 62
		encoded = string(base62Alphabet[remainder]) + encoded
		num /= 62
	}
	return encoded
}

func base62Dec(str string) (int64, error) {
	if len(str) == 0 {
		return 0, errors.New("empty string")
	}

	var res int64

	for _, r := range str {
		res *= 62
		val := int64(strings.Index(base62Alphabet, string(r)))
		res += val
	}
	return res, nil
}

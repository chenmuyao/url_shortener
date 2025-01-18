package service

import "context"

type UrlShortenerSvc interface {
	// Shorten takes the full url and returns a short url
	Shorten(ctx context.Context, fullUrl string) (string, error)
	// GetFull takes a shortID and returns the full url
	GetFull(ctx context.Context, shortID string) (string, error)
}

type urlShortenerSvc struct{}

// GetFull implements UrlShortenerSvc.
func (u *urlShortenerSvc) GetFull(ctx context.Context, shortID string) (string, error) {
	return "http://www.example.com", nil
}

// Shorten implements UrlShortenerSvc.
func (u *urlShortenerSvc) Shorten(ctx context.Context, fullUrl string) (string, error) {
	return "sdf123", nil
}

func NewUrlShortenerSvc() UrlShortenerSvc {
	return &urlShortenerSvc{}
}

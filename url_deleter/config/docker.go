//go:build docker

package config

import "time"

var App = AppConfig{
	DBAddr:      "postgres://postgres:postgres@postgres:5432/url_shortener?sslmode=disable",
	Schedule:    "@daily",
	DataTimeout: 7 * 24 * time.Hour,
	JobTimeout:  24 * time.Hour,
}

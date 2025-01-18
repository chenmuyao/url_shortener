//go:build !docker

package config

import "time"

var App = AppConfig{
	DBAddr:      "postgres://postgres:postgres@localhost:25432/url_shortener?sslmode=disable",
	Schedule:    "@every 10s",
	DataTimeout: time.Second * 10,
	JobTimeout:  time.Second * 10,
}

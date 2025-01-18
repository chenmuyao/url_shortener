//go:build !docker

package config

var DB = DBConfig{
	Addr:             "postgres://postgres:postgres@localhost:25432/url_shortener?sslmode=disable",
	MaxOpenDbConn:    100,
	MaxIdleDbConn:    100,
	MaxDbLifetimeSec: 10 * 60,
}

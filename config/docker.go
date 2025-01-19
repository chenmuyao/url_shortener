//go:build docker

package config

var DB = DBConfig{
	Addr:             "postgres://postgres:postgres@postgres:5432/url_shortener?sslmode=disable",
	MaxOpenDbConn:    100,
	MaxIdleDbConn:    100,
	MaxDbLifetimeSec: 10 * 60,
}

var Redis = RedisConfig{
	Addr: "redis:6379",
}

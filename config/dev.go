//go:build !docker

package config

var DB = DBConfig{
	Addr:             "postgres://postgres:postgres@localhost:25432/url_shortener?sslmode=disable",
	MaxOpenDbConn:    100,
	MaxIdleDbConn:    100,
	MaxDbLifetimeSec: 10 * 60,
}

var Redis = RedisConfig{
	Addr: "localhost:26379",
}

var App = Application{
	BaseURL: "http://localhost:3000",
}

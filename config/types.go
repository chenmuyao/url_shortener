package config

type DBConfig struct {
	Addr             string
	MaxOpenDbConn    int
	MaxIdleDbConn    int
	MaxDbLifetimeSec int
}

type RedisConfig struct {
	Addr string
}

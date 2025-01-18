package config

type DBConfig struct {
	Addr             string
	MaxOpenDbConn    int
	MaxIdleDbConn    int
	MaxDbLifetimeSec int
}

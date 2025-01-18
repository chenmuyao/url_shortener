package config

import "time"

type AppConfig struct {
	DBAddr string
	// In the cron format, see robfig/cron
	Schedule    string
	DataTimeout time.Duration
	JobTimeout  time.Duration
}

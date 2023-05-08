package config

import "time"

type Config struct {
	Addr          string
	ApiPath       string
	AdminPath     string
	AdminUsername string
	LifeSpan      int
	Key           string
	StartTime     time.Time
}

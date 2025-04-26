package cfg

import (
	"encoding/json"
	"log/slog"
	"os"
	"sync"

	"github.com/caarlos0/env/v11"
)

var (
	once sync.Once
	cfg  config
)

type config struct {
	LogLevel string      `json:"log_level"`
	Redis    redisConfig `json:"redis"`
	Postgres pgConfig    `json:"postgres"`
	Secrets  secrets
}

type redisConfig struct {
	Address string `json:"address"` // host and port ( localhost:6379 )
	Db      int    `json:"db"`
}

type pgConfig struct {
	Username string `json:"username"`
	Address  string `json:"address"` // host and port ( localhost:6379 )
	DbName   string `json:"db_name"`
	Schema   string `json:"schema"`
}

type secrets struct {
	RedisPassword    string `env:"REDIS_PASSWORD"`
	PostgresPassword string `env:"PG_PASSWORD"`
}

func GetLogLevel(logLevel string) slog.Level {
	switch logLevel {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
func Cfg() config {
	once.Do(func() {

		configFile, err := os.Open("config.json")
		if err != nil {
			panic(err.Error())
		}
		defer configFile.Close()
		if err := json.NewDecoder(configFile).Decode(&cfg); err != nil {
			panic(err.Error())
		}

		if err := env.Parse(&cfg.Secrets); err != nil {
			panic(err.Error())
		}

	})
	return cfg
}

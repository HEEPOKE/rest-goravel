package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("database", map[string]any{
		"default": config.Env("DB_CONNECTION", "postgresql"),

		"connections": map[string]any{
			"mysql": map[string]any{
				"driver":   "mysql",
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", 3306),
				"database": config.Env("DB_DATABASE", "forge"),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"charset":  "utf8mb4",
				"loc":      "Local",
				"prefix":   "",
				"singular": false,
			},
			"postgresql": map[string]any{
				"driver":   "postgresql",
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", 5432),
				"database": config.Env("DB_DATABASE", ""),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"sslmode":  "disable",
				"timezone": "UTC",
				"prefix":   "",
				"singular": false,
			},
			"sqlite": map[string]any{
				"driver":   "sqlite",
				"database": config.Env("DB_DATABASE", "forge"),
				"prefix":   "",
				"singular": false,
			},
			"sqlserver": map[string]any{
				"driver":   "sqlserver",
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", 1433),
				"database": config.Env("DB_DATABASE", "forge"),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"charset":  "utf8mb4",
				"prefix":   "",
				"singular": false,
			},
		},
		"pool": map[string]any{
			"max_idle_conns":    10,
			"max_open_conns":    100,
			"conn_max_idletime": 3600,
			"conn_max_lifetime": 3600,
		},
		"migrations": "migrations",
		"redis": map[string]any{
			"default": map[string]any{
				"host":     config.Env("REDIS_HOST", ""),
				"password": config.Env("REDIS_PASSWORD", ""),
				"port":     config.Env("REDIS_PORT", 6379),
				"database": config.Env("REDIS_DB", 0),
			},
		},
	})
}

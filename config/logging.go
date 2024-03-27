package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("logging", map[string]any{
		"default": config.Env("LOG_CHANNEL", "stack"),
		"channels": map[string]any{
			"stack": map[string]any{
				"driver":   "stack",
				"channels": []string{"daily"},
			},
			"single": map[string]any{
				"driver": "single",
				"path":   "storage/logs/goravel.log",
				"level":  config.Env("LOG_LEVEL", "debug"),
				"print":  true,
			},
			"daily": map[string]any{
				"driver": "daily",
				"path":   "storage/logs/goravel.log",
				"level":  config.Env("LOG_LEVEL", "debug"),
				"days":   7,
				"print":  true,
			},
		},
	})
}

package config

import "github.com/goravel/framework/facades"

func init() {
	config := facades.Config()
	config.Add("mail", map[string]any{
		"host": config.Env("MAIL_HOST", ""),
		"port": config.Env("MAIL_PORT", 587),
		"from": map[string]any{
			"address": config.Env("MAIL_FROM_ADDRESS", "hello@example.com"),
			"name":    config.Env("MAIL_FROM_NAME", "Example"),
		},
		"username": config.Env("MAIL_USERNAME"),
		"password": config.Env("MAIL_PASSWORD"),
	})
}

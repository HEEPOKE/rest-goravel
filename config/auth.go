package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("auth", map[string]any{
		"defaults": map[string]any{
			"guard": "user",
		},
		"guards": map[string]any{
			"user": map[string]any{
				"driver": "jwt",
			},
		},
	})
}

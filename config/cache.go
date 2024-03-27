package config

import (
	"github.com/goravel/framework/contracts/cache"
	"github.com/goravel/framework/facades"
	redisFacades "github.com/goravel/redis/facades"
)

func init() {
	config := facades.Config()
	config.Add("cache", map[string]any{
		"default": config.Env("CACHE_STORE", "memory"),
		"stores": map[string]any{
			"redis": map[string]any{
				"driver":     "custom",
				"connection": "default",
				"via": func() (cache.Driver, error) {
					return redisFacades.Redis("redis"), nil
				},
			},
		},
		"prefix": config.GetString("APP_NAME", "goravel") + "_cache",
	})
}

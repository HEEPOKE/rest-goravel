package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("grpc", map[string]any{
		"host":    config.Env("GRPC_HOST", ""),
		"port":    config.Env("GRPC_PORT", ""),
		"clients": map[string]any{
			//"user": map[string]any{
			//	"host":         config.Env("GRPC_USER_HOST", ""),
			//	"port":         config.Env("GRPC_USER_PORT", ""),
			//	"interceptors": []string{},
			//},
		},
	})
}

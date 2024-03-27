package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	fiberFacades "github.com/goravel/fiber/facades"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("http", map[string]any{
		"default": "fiber",

		"drivers": map[string]any{
			"fiber": map[string]any{
				"prefork":      false,
				"body_limit":   4096,
				"header_limit": 4096,
				"route": func() (route.Route, error) {
					return fiberFacades.Route("fiber"), nil
				},
				"template": func() (fiber.Views, error) {
					return html.New("./resources/views", ".tmpl"), nil
				},
			},
		},
		"url":  config.Env("APP_URL", "http://localhost"),
		"host": config.Env("APP_HOST", "127.0.0.1"),
		"port": config.Env("APP_PORT", "3000"),
		"tls": map[string]any{
			"host": config.Env("APP_HOST", "127.0.0.1"),
			"port": config.Env("APP_PORT", "3000"),
			"ssl": map[string]any{
				"cert": "",
				"key":  "",
			},
		},
	})
}

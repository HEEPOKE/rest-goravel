package main

import (
	"github.com/goravel/framework/facades"

	"goravel/bootstrap"
)

func main() {
	bootstrap.Boot()

	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route run error: %v", err)
		}
	}()

	select {}
}

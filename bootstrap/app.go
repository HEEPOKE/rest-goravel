package bootstrap

import (
	"github.com/goravel/framework/foundation"

	"goravel/config"
)

func Boot() {
	app := foundation.NewApplication()

	app.Boot()

	config.Boot()
}

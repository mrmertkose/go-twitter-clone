package bootstrap

import (
	"twitterClone/pkg/config"
	"twitterClone/pkg/database"
	"twitterClone/pkg/routing"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}

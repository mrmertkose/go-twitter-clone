package bootstrap

import (
	"twitterClone/pkg/config"
	"twitterClone/pkg/database"
	"twitterClone/pkg/routing"
)

func Serve() {
	config.Set()

	database.Connect()

	//database.DB.AutoMigrate(&models.User{})

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}

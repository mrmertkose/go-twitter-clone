package bootstrap

import "twitterClone/pkg/config"

func Serve() {
	config.Set()
	
	//database.Connect()
	//
	//routing.Init()
	//
	//sessions.Start(routing.GetRouter())
	//
	//static.LoadStatic(routing.GetRouter())
	//
	//html.LoadHtml(routing.GetRouter())
	//
	//routing.RegisterRoutes()
	//
	//routing.Serve()
}

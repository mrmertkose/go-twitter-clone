package routes

import (
	"github.com/gin-gonic/gin"
	"twitterClone/internal/modules/user/controllers"
)

func Routes(router *gin.Engine) {
	userController := controllers.New()

	router.GET("/login", userController.Login)
	router.GET("/register", userController.Register)
}

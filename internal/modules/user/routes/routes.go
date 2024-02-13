package routes

import (
	"github.com/gin-gonic/gin"
	"twitterClone/internal/modules/user/controllers"
)

func Routes(router *gin.Engine) {
	userController := controllers.New()

	router.POST("/login", userController.Login)
	router.POST("/register", userController.Register)
}

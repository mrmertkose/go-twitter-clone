package routes

import (
	"github.com/gin-gonic/gin"
	userRoutes "twitterClone/internal/modules/user/routes"
)

func RegisterRoutes(router *gin.Engine) {
	userRoutes.Routes(router)
}

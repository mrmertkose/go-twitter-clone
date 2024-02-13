package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"twitterClone/internal/modules/user/request"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (controller Controller) Login(c *gin.Context) {
	var loginRequest request.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(loginRequest)
}

func (controller Controller) Register(context *gin.Context) {

}

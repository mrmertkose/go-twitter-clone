package controllers

import "github.com/gin-gonic/gin"

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c Controller) Login(context *gin.Context) {

}

func (c Controller) Register(context *gin.Context) {

}

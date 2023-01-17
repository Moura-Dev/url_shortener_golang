package controllers

import (
	"github.com/gin-gonic/gin"
)

type IController interface {
	CreateGoly(ctx *gin.Context)
	GetAllGolies(ctx *gin.Context)
	GetGoly(ctx *gin.Context)
	GetGolyUuid(ctx *gin.Context)
	DeleteGoly(ctx *gin.Context)
	UpdateGoly(ctx *gin.Context)
	Redirect(ctx *gin.Context)
}

type Controller struct {
}

func NewController() Controller {
	return Controller{}
}

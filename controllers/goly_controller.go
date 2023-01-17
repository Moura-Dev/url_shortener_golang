package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"url_shortner/models"
	"url_shortner/repository"
	"url_shortner/utils"
)

func (c *Controller) CreateGoly(ctx *gin.Context) {
	var goly models.Goly
	if err := ctx.ShouldBindJSON(&goly); err != nil {
		ctx.JSON(400, gin.H{
			"error": "failed json binding",
		})
		return
	}
	if goly.Random {
		goly.Goly = utils.RandomGoly(8)
	}
	goly, err := repository.CreateGoly(goly)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "failed to create goly",
		})
		return
	}
	ctx.JSON(201, goly)
}

func (c *Controller) GetAllGolies(ctx *gin.Context) {
	golies, err := repository.GetAllGolies()
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "could not get golies",
		})
		return
	}
	ctx.JSON(200, golies)
}

func (c *Controller) GetGoly(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	goly, err := repository.GetGoly(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "could not find goly",
		})
		return
	}
	ctx.JSON(200, goly)
}

func (c *Controller) GetGolyUuid(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	goly, err := repository.GetGolyUrl(uuid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "could not find goly",
		})
		return
	}
	ctx.JSON(200, goly)
}

func (c *Controller) DeleteGoly(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	err = repository.DeleteGoly(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Delete failed",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Goly deleted successfully",
	})
}

func (c *Controller) UpdateGoly(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var goly models.Goly
	if err := ctx.ShouldBindJSON(&goly); err != nil {
		ctx.JSON(400, gin.H{
			"error": "could not update goly",
		})
		return
	}
	goly.ID = id
	goly, err = repository.UpdateGoly(goly)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, goly)
}

func (c *Controller) Redirect(ctx *gin.Context) {
	redirect := ctx.Param("redirect")
	goly, err := repository.GetGolyUrl(redirect)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "could not redirect",
		})
		return
	}
	goly.Count += 1
	ctx.Redirect(301, goly.RedirectURL)
}

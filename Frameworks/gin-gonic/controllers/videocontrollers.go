package controllers

import (
	"GIN-GONIC/entity"
	"GIN-GONIC/services"

	"github.com/gin-gonic/gin"
)


type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type Controller struct{
	service services.VideoService
}

func New(Service services.VideoService) Controller{
	return Controller{
		service:Service,
	}
}
func (c *Controller)FindAll() []entity.Video{
	return c.service.FindAll()
}

func (c *Controller) Save(ctx *gin.Context)entity.Video{
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
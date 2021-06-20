package main

import (
	"GIN-GONIC/controllers"
	"GIN-GONIC/services"

	"github.com/gin-gonic/gin"
)
var(
	videoService services.VideoService =services.New()
	VideoController controllers.Controller = controllers.New(videoService)
)


func main() {
	server := gin.Default()
	server.GET("/videos",func(ctx *gin.Context){
		ctx.JSON(200,VideoController.FindAll())
		})
	server.POST("/videos",func(ctx *gin.Context){
			ctx.JSON(200,VideoController.Save(ctx))
			})

	

	server.Run(":8080")
}
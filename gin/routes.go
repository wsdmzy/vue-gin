package main

import (
	"github.com/gin-gonic/gin"
	"ziogie.top/gin/controller"
	"ziogie.top/gin/middleware"
)

func CollectRoute(r *gin.Engine)  *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return  r
}
package main

import (
	"github.com/gin-gonic/gin"
	"ziogie.top/gin/controller"
	"ziogie.top/gin/middleware"
)

func CollectRoute(r *gin.Engine)  *gin.Engine {
	r.Use(middleware.CORSMiddleware(),middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	//分类页路由
	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("",categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)
	//categoryRoutes.PATCH() 修改     put替换


	return  r
}
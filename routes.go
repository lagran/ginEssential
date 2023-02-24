package main

import (
	"funboy.top/ginessential/controller"
	"funboy.top/ginessential/middlewarre"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middlewarre.AuthMiddleware(), controller.Info)
	return r
}

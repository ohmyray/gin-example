package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ohmyray/gin-example/controller"
	"github.com/ohmyray/gin-example/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}

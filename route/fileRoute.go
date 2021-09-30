package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ohmyray/gin-example/controller"
	// "github.com/ohmyray/gin-example/middleware"
)

func InstallFileRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/file/upload", controller.Upload)

	return r
}
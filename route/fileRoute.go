package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ohmyray/gin-example/controller"
	// "github.com/ohmyray/gin-example/middleware"
)

func InstallFileRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/file/upload", controller.Upload)
	r.GET("/api/file/upload", controller.FindUpload)
	r.GET("/api/file/upload/:id", controller.FindUploadById)
	r.DELETE("/api/file/upload/:id", controller.DeleteUploadById)

	return r
}
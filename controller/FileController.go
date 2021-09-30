package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ohmyray/gin-example/model"
	"github.com/ohmyray/gin-example/response"
)

func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.Fail(ctx, nil, "未选择文件")
		return
	}

	var bindFile model.File

	if err = ctx.ShouldBind(&bindFile); err != nil {
		response.Fail(ctx, nil, "请求参数错误")
		return
	}
	
	// fmt.Println(bindFile.Name, bindFile.FilePath)
	fmt.Println(file.Filename)
	// fmt.Println(file)
	ctx.SaveUploadedFile(file, file.Filename)
	ctx.JSON(http.StatusOK, gin.H{
		"code": "200",
	})


}
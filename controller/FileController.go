package controller

import (
	"fmt"
	"github.com/ohmyray/gin-example/common"
	"github.com/ohmyray/gin-example/dto"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/ohmyray/gin-example/model"
	"github.com/ohmyray/gin-example/response"
)

const BASE_PATH = "upload/"

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

	path := filepath.Base(file.Filename)
	saveFilePath := BASE_PATH + path
	ctx.SaveUploadedFile(file, saveFilePath)
	fmt.Println(file.Filename,path)
	db := common.GetDB()

	newFile := model.File{
		Name:      "ohmyray",
		FilePath:  saveFilePath,
	}

	db.Create(&newFile)
	response.Success(ctx, gin.H{"file": dto.TransformToFileDto(newFile)}, "上传成功")


}
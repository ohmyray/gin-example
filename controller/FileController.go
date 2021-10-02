package controller

import (
	"fmt"
	"github.com/ohmyray/gin-example/common"
	"github.com/ohmyray/gin-example/dto"
	"os"
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

func FindUpload(ctx *gin.Context) {
	var bindUser model.User
	var bindFiles []model.File
	ctx.ShouldBind(&bindUser)
	db := common.GetDB()
	db.Where("name = ?", bindUser.Name).Find(&bindFiles)

	//if len(bindFiles) != 0 {
	//	response.Success()
	//	return
	//}

	var fileSlice []dto.FileDto
	for _, file := range bindFiles {
		fileSlice = append(fileSlice, dto.FileDto{
			Id: int(file.ID),
			Name:       file.Name,
			FilePath:   file.FilePath,
			UploadTime: file.CreatedAt.String(),
		})
	}
	response.Success(ctx, gin.H{ "fileList": fileSlice },"查询成功")
}

func FindUploadById(ctx *gin.Context){
	id := ctx.Param("id")

	var queryFile model.File
	db := common.GetDB()
	db.First(&queryFile, id)

	if queryFile.ID != 0 {
		response.Fail(ctx, nil, "未查询指定到文件")
	}

	response.Success(ctx,gin.H{"file": queryFile},"查询成功")
}

func DeleteUploadById(ctx *gin.Context )  {
	id := ctx.Param("id")

	var queryFile model.File

	db := common.GetDB()
	db.Where("id = ?", id).First(&queryFile)

	if queryFile.ID == 0 {
		response.Fail(ctx, nil, "未查询指定到文件")
		return
	}
	err := os.Remove(queryFile.FilePath)
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}
	db.Delete(&queryFile)
	response.Success(ctx, nil, "删除成功")
}
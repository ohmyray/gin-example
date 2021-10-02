package dto

import "github.com/ohmyray/gin-example/model"

type FileDto struct {
	Id        int `json:"id"`
	Name      string `json:"name"`
	FilePath  string `json:"filePath"`
	UploadTime  string `json:"uploadTime"`
}

func TransformToFileDto(file model.File) FileDto {
	return FileDto{
		Id: int(file.ID),
		Name:       file.Name,
		FilePath:   file.FilePath,
		UploadTime: file.CreatedAt.String(),
	}
}

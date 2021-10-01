package dto

import "github.com/ohmyray/gin-example/model"

type FileDto struct {
Name      string `json:"name"`
FilePath string `json:"filePath"`
}

func TransformToFileDto(file model.File) FileDto {
	return FileDto{
		Name:      file.Name,
		FilePath:file.FilePath,
	}
}

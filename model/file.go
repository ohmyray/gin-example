package model

import (
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	Name      string `form:"name" json:"name" gorm:"type:varchar(20);not null"`
	FilePath string `form:"filePath" json:"filePath" gorm:"type:varchar(255);not null"`
	// File  string `form:"file" json:"file" gorm:"type:varchar(255);not null"`
}

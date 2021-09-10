package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `form:"name" json:"name" gorm:"type:varchar(20);not null"`
	Telephone string `form:"telephone" json:"telephone" gorm:"type:varchar(11);not null"`
	Password  string `form:"password" json:"password" gorm:"type:varchar(255);not null"`
}

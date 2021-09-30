package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ohmyray/gin-example/model"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB(viper *viper.Viper) *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	sourceName := viper.GetString("datasource.sourceName")

	args := ""
	switch(driverName){
		case "mysql":
			host := viper.GetString("datasource.host")
			port := viper.GetString("datasource.port")
			database := viper.GetString("datasource.database")
			username := viper.GetString("datasource.username")
			password := viper.GetString("datasource.password")
			charset := viper.GetString("datasource.charset")
			args = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
				username, password, host, port, database, charset)
		case "sqlite3":
			args = sourceName + ".db"
	}

	db, err := gorm.Open(driverName, args)

	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/ohmyray/gin-example/common"
	"github.com/ohmyray/gin-example/middleware"
	"github.com/ohmyray/gin-example/route"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var config *viper.Viper

func main() {
	// 1.读取配置文件
	config = initConfigure()
	// 2.通过配置连接数据库
	db := common.InitDB(config)
	defer db.Close()

	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r = route.CollectRoute(r)
	r = route.InstallFileRoute(r)

	r.GET("/getConfig", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"config": config.AllSettings(),
		})
	})
	port := config.GetString("server.port")

	if port == "" {
		panic(r.Run())
	}

	fmt.Println("listen and serve on 0.0.0.0:" + port)

	panic(r.Run(":" + port))
}

func initConfigure() *viper.Viper {
	v := viper.New()
	v.SetConfigName("application") // 设置文件名称（无后缀）
	v.SetConfigType("yaml")        // 设置后缀名 {"1.6以后的版本可以不设置该后缀"}
	v.AddConfigPath("./conf")      // 设置文件所在路径
	v.Set("verbose", true)         // 设置默认参数

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(" Config file not found; ignore error if desired")
		} else {
			panic("Config file was found but another error was produced")
		}
	}
	// 监控配置和重新获取配置
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name, e)
	})
	return v
}

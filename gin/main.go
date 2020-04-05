package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"ziogie.top/gin/common"
)


func main() {
	InitConfig()
	db := common.GetDB()
	defer db.Close()
	r := gin.Default()

	r = CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":"+port))
	}
	//默认 监听并在 0.0.0.0:8080 上启动服务
	panic(r.Run(":8080"))
}

//读取配置
func InitConfig()  {
	workDir,_ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}






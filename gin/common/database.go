package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"ziogie.top/gin/model"
)

//连接数据库
func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db,err := gorm.Open(driverName, args)
	if err != nil {
		panic("faild to connect database, err" + err.Error())
	}
	//自动创建数据表
	db.AutoMigrate(&model.User{})
	//fmt.Println("++")
	return  db
}

func GetDB() *gorm.DB  {
	return InitDB()
}
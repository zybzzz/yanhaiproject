/*
	包 core 包含服务端核心组件
 */
package core

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
)

var ApplicationConfig *viper.Viper

/**
 * @author ZhangYiBo
 * @description 初始化viper读取配置参数
 * @date 11:43 下午 2021/9/15
 * @param
 * @return
 **/
func init() {
	ApplicationConfig = viper.New()
	ApplicationConfig.SetConfigName("ApplicationConfig")
	ApplicationConfig.SetConfigType("yml")
	ApplicationConfig.AddConfigPath(".")
	//在这里可以设置默认值
	if err := ApplicationConfig.ReadInConfig(); err != nil{
		log.Printf("read config failed: %v\n", err)
		log.Panic("初始化viper失败！程序崩溃")
	}
	log.Info("初始化viper成功")
}


var DB *gorm.DB
var err error

/**
 * @author ZhangYiBo
 * @description  初始化数据库
 * @date 11:44 下午 2021/9/15
 * @param
 * @return
 **/
func init()  {
	user := ApplicationConfig.GetString("database.user")
	password := ApplicationConfig.GetString("database.password")
	server := ApplicationConfig.GetString("database.server")
	port := ApplicationConfig.GetString("database.port")
	dbName := ApplicationConfig.GetString("database.dbName")
	dsn := user + ":" + password + "@tcp(" + server + ":" + port + ")/" + dbName +"?charset=utf8&parseTime=True"
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("连接数据库失败: %v\n", err)
		log.Panic("连接数据库失败，程序崩溃")
	}
	log.Info("数据库连接成功")
}

const URL_PREFIX = "http://47.96.230.189:9225/pics/"
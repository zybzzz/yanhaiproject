/*
	包 core 包含服务端核心组件
 */
package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

//FIXME 抽取配置信息
func init()  {
	dsn := "root:password@tcp(localhost:3306)/Dbname?charset=utf8&parseTime=True"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//FIXME 更改为log打印日志
		fmt.Println("连接数据库成功")
	}
}
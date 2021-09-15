/*
	包main 程序执行从此开始
 */
package main

import (
	"github.com/gin-gonic/gin"
	"yanhaiproject/router"
	_ "yanhaiproject/core"
)

func main() {
	engine := gin.Default()

	router.ApiV1RouterInit(engine)

	//FIXME 补充端口信息
	engine.Run()
}

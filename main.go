/*
	包 main 程序执行从此开始
 */
package main

import (
	"github.com/gin-gonic/gin"
	"yanhaiproject/core"
	"yanhaiproject/middlewares"
	"yanhaiproject/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	engine := gin.Default()
	log.Info("nihao")
	engine.Use(middlewares.Cors())
	router.ApiV1RouterInit(engine)

	//FIXME 补充端口信息
	runport := ":" + core.ApplicationConfig.GetString("port")
	engine.Run(runport)
}

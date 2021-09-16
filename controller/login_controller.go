/*
	包 controller 实现控制层相关处理
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yanhaiproject/core"
	"yanhaiproject/model"
	log "github.com/sirupsen/logrus"
)

type LoginController struct {
}

func (con LoginController) Login(context *gin.Context)  {
	var user model.User
	bodyJson := make(map[string]interface{})
	context.BindJSON(&bodyJson)
	log.Info(bodyJson)
	core.DB.Where("user_id = ?", bodyJson["userId"]).First(&user)
	log.Info("查询出的user..." )
	log.Info(user)
	//TODO 暂时没有对数据库user为空的情况进行判断
	if user.Password == bodyJson["password"] {
		context.JSON(http.StatusOK, gin.H{
			"status": "success",
			"userId": user.UserId,
		})
	}else {
		context.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
	}
}

func (con LoginController) Register(context *gin.Context)  {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil{
		context.String(http.StatusBadRequest, err.Error())
	}else {
		log.Info("register 请求中的user")
		log.Info(user)
		result := core.DB.Omit("Gender", "PortraitId", "AttentionId").Create(&user)
		if result.Error != nil {
			log.Info("注册数据插入成功")
			context.JSON(http.StatusOK, gin.H{
				"status": "success",
				"userId": user.UserId,
			})
		}else {
			log.Error("插入数据出现错误  " + result.Error.Error())
		}
	}
}
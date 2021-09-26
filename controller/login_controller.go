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
	userId := context.Query("userId")
	password := context.Query("password")
	core.DB.Where("user_id = ?", userId).First(&user)
	log.Info("查询出的user..." )
	log.Info(user)
	//TODO 暂时没有对数据库user为空的情况进行判断
	if user.Password == password {
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
		//FIXME 此处忽略了一堆属性
		user.PortraitId = 1
		result := core.DB.Omit( "AttentionId").Create(&user)
		if result.Error == nil {
			log.Info("注册数据插入成功")
			context.JSON(http.StatusOK, gin.H{
				"status": "success",
				"userId": user.UserId,
			})
		}else {
			context.JSON(http.StatusOK, gin.H{
				"status": "fail",
			})
			log.Error("插入数据出现错误  " + result.Error.Error())
		}
	}
}
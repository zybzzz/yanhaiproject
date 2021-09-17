package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yanhaiproject/core"
	"yanhaiproject/model"
	log "github.com/sirupsen/logrus"
)

type CommentController struct {
}

func (con CommentController) Comment(context *gin.Context)  {
	var comment model.Comment
	if err := context.ShouldBindJSON(&comment); err != nil{
		log.Error(err)
		context.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
		return
	}
	log.Info("接收到的comment是")
	log.Info(comment)
	result := core.DB.Create(&comment)
	if result.Error != nil {
		log.Error(result.Error)
		context.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
	}else {
		context.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

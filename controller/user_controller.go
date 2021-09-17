package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yanhaiproject/service"
)

type UserController struct {
}

func (con UserController) GetMyMessage(context *gin.Context)  {
	retMess, isSuccess := service.UserService{}.GetUserMessByUserId(context.Query("userId"))
	if isSuccess {
		context.String(http.StatusOK, retMess)
	}else {
		context.JSON(http.StatusOK,gin.H{
			"status": "fail",
		})
	}
}

func (con UserController) ChangeMyMess(context *gin.Context)  {
	isSuccess := service.UserService{}.ChangeMyMessByUserId(context.Query("userId"),context)
	if isSuccess {
		context.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}else {
		context.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
	}
}

func (con UserController) GetMyReleaseList(context *gin.Context)  {

}

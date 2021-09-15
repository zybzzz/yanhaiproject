package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SchoolGroupController struct {
}

func (con SchoolGroupController) GetGroupList(context *gin.Context)  {

}

func (con SchoolGroupController) GetGroupDetail(context *gin.Context)  {
	groupId := context.Param("groupId")
	context.JSON(http.StatusOK, gin.H{
		"groupId": groupId,
	})
}
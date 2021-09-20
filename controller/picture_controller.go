package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yanhaiproject/service"
)

type PictureController struct {
}

func (con PictureController) UploadPicture(context *gin.Context)  {

	_, isSuccess, ids := service.StorePictures(context)
	if isSuccess == true {
		context.JSON(http.StatusOK,gin.H{
			"status":"success",
			"picids":ids,
		})
	}else {
		context.JSON(http.StatusOK,gin.H{
			"status":"fail",
		})
	}

}

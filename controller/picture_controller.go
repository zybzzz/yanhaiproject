package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yanhaiproject/service"
	log "github.com/sirupsen/logrus"
)

type PictureController struct {
}

func (con PictureController) UploadPicture(context *gin.Context)  {

	mess, isSuccess, ids, urls := service.StorePictures(context)
	if isSuccess == true {
		context.JSON(http.StatusOK,gin.H{
			"status":"success",
			"picids":ids,
			"picurls":urls,
		})
	}else {
		log.Debug(mess)
		context.JSON(http.StatusOK,gin.H{
			"status":"fail",
			"mess":mess,
		})
	}

}

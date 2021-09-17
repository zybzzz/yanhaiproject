/*
	包 service 涉及服务层相关代码
 */
package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"yanhaiproject/core"
	"yanhaiproject/model"
)

type UserService struct {
}

type attention struct {
	School            string `json:"school"`
	SchoolPortraitURL string `json:"schoolPortraitURL"`
}

type userMess struct {
	Status            string `json:"status"`
	UserID            string `json:"userId"`
	Gender            string `json:"gender"`
	Nickname          string `json:"nickname"`
	PortraitURL       string `json:"portraitURL"`
	AttentionListSize int    `json:"attentionListSize"`
	AttentionList     []attention `json:"attentionList"`
	School string `json:"school"`
	Major  string `json:"major"`
}

/**
 * @author ZhangYiBo
 * @description  通过用户id获取用户信息
 * @date 11:48 上午 2021/9/16
 * @param userId 用户id
 * @return json字符串或错误提示信息 操作是否成功
 **/
func (UserService) GetUserMessByUserId(userId string) (string, bool) {
	var user model.User
	result := core.DB.First(&user, userId)
	if result.RowsAffected == 0 {
		return "没有该用户",false
	}
	var retmess userMess
	retmess.Status = "success"
	retmess.Nickname = user.Nickname
	retmess.UserID = user.UserId
	retmess.Gender = user.Gender
	retmess.School = user.School
	retmess.Major = user.Major
	retmess.PortraitURL = PictureService{}.PicIdToURL(user.PortraitId)
	attentionList := GroupService{}.getAttentionList("|")
	retmess.AttentionListSize = len(attentionList)
	retmess.AttentionList = attentionList
	bytejson, err := json.Marshal(retmess)
	if err != nil {
		log.Info(err.Error())
	}
	log.Info("返回的字符串是" + string(bytejson))
	return string(bytejson),true
}

func (UserService) ChangeMyMessByUserId(userId string, context *gin.Context) (isSuccess bool) {
	var user model.User
	result := core.DB.First(&user, userId)
	if result.Error != nil{
		isSuccess = false
		return
	}
	err := context.ShouldBindJSON(&user)
	if err != nil {
		isSuccess = false
		return
	}
	log.Info("当前更新的user是")
	log.Info(user)
	result = core.DB.Save(&user)
	if result.Error != nil {
		isSuccess = false
		return
	}
	return true
}

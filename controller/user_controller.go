package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"yanhaiproject/core"
	"yanhaiproject/model"
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
	userId := context.Query("userId")
	log.Info("获取到的用户id")
	log.Info(userId)
	//获取用户的头像数据
	var user model.User
	core.DB.First(&user, userId)
	userPortrait := service.PictureService{}.PicIdToURL(user.PortraitId)
	//从用户关注的专业和关注的学校给用户推荐帖子
	var topics []model.Topic
	result := core.DB.Where("creator = ?", userId).Find(&topics)
	if result.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
	}
	var retTopics = make([]retTopic, len(topics))
	for index , topic := range topics{
		retTopics[index].TopicId = topic.TopicId
		retTopics[index].Creator = topic.Creator
		var creator model.User
		core.DB.First(&creator, topic.Creator)
		retTopics[index].Nickname = creator.Nickname
		retTopics[index].Title = topic.Title
		retTopics[index].Content = topic.Content
		retTopics[index].CreateTime = topic.CreateAt.Format("2006-01-02 15:04:05")
		retTopics[index].ThumpUp = topic.ThumpUp
		//FIXME 暂时写死 在评论的时候直接设置字段自增 等待优化
		retTopics[index].RecommendNum = 20
		retTopics[index].Portrait = userPortrait
		retTopics[index].Tag = strings.Split(topic.Tag, "|")
		retTopics[index].TopicPictures = service.PictureService{}.PicIdsToURL(topic.PicId)
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"size": len(retTopics),
		"topicList": retTopics,
	})
}

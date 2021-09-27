package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"yanhaiproject/core"
	"yanhaiproject/model"
	log "github.com/sirupsen/logrus"
	"yanhaiproject/service"
)

type IndexPageController struct {
}

type retTopic struct {
	TopicId  int `json:"topicId"`
	Creator  string `json:"creator"`
	Nickname string `json:"nickname"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	CreateTime string `json:"createTime"`
	ThumpUp  int `json:"thumpUp"`
	Tag      []string `json:"tag"`
	RecommendNum int `json:"recommendNum"`
	TopicPictures []string `json:"topicPictures"`
	Portrait string `json:"portrait"`
}

func (con IndexPageController) GetRecommendList(context *gin.Context)  {
	var user model.User
	user.UserId = context.Query("userId")
	core.DB.First(&user, user.UserId)
	//从用户关注的专业给用户推荐帖子
	var topics []model.Topic
	result := core.DB.Where("major = ?", user.Major).Find(&topics)
	if result.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
	}
	var retTopics = make([]retTopic, len(topics))
	for index , topic := range topics{
		retTopics[index].TopicId = topic.TopicId
		retTopics[index].Creator = topic.Creator
		retTopics[index].Title = topic.Title
		retTopics[index].Content = topic.Content
		retTopics[index].CreateTime = topic.CreateAt.Format("2006-01-02 15:04:05")
		retTopics[index].ThumpUp = topic.ThumpUp
		retTopics[index].Tag = strings.Split(topic.Tag, "|")
		//FIXME 暂时写死 在评论的时候直接设置字段自增 等待优化
		retTopics[index].RecommendNum = 20
		var creator model.User
		core.DB.First(&creator, topic.Creator)
		retTopics[index].Nickname = creator.Nickname
		retTopics[index].TopicPictures = service.PictureService{}.PicIdsToURL(topic.PicId)
		retTopics[index].Portrait = service.PictureService{}.PicIdToURL(creator.PortraitId)
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"size": len(retTopics),
		"topicList": retTopics,
	})
}

func (con IndexPageController) GetAttentionList(context *gin.Context)  {
	var user model.User
	user.UserId = context.Query("userId")
	core.DB.First(&user, user.UserId)
	attentions := strings.Split(user.AttentionId, "|")
	attentionIds := make([]int, len(attentions))
	for index, attentionId := range attentions{
		attentionIds[index], _ = strconv.Atoi(attentionId)
	}
	log.Info("获取到的关注id")
	log.Info(attentionIds)
	//从用户关注的专业和关注的学校给用户推荐帖子
	var topics []model.Topic
	result := core.DB.Where(map[string]interface{}{"major":user.Major, "group_id":attentionIds}).Find(&topics)
	if result.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
	}
	var retTopics = make([]retTopic, len(topics))
	for index , topic := range topics{
		retTopics[index].TopicId = topic.TopicId
		retTopics[index].Creator = topic.Creator
		retTopics[index].Title = topic.Title
		retTopics[index].Content = topic.Content
		retTopics[index].CreateTime = topic.CreateAt.Format("2006-01-02 15:04:05")
		retTopics[index].ThumpUp = topic.ThumpUp
		//FIXME 暂时写死 在评论的时候直接设置字段自增 等待优化
		retTopics[index].RecommendNum = 20
		retTopics[index].Tag = strings.Split(topic.Tag, "|")
		var creator model.User
		core.DB.First(&creator, topic.Creator)
		retTopics[index].Nickname = creator.Nickname
		retTopics[index].TopicPictures = service.PictureService{}.PicIdsToURL(topic.PicId)
		retTopics[index].Portrait = service.PictureService{}.PicIdToURL(creator.PortraitId)
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"size": len(retTopics),
		"topicList": retTopics,
	})
}

package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"yanhaiproject/core"
	"yanhaiproject/model"
)

type TopicController struct {
}

type retComment struct {
	Nickname    string `json:"nickname"`
	PortraitURL string `json:"portraitURL"`
	CreateTime  string `json:"createTime"`
	Content     string `json:"content"`
	ThumpUp     int    `json:"thumpUp"`
}

type topicDetail struct {
	Status          string   `json:"status"`
	PortraitURL     string   `json:"portraitURL"`
	Nickname        string   `json:"nickname"`
	CreateTime      string   `json:"createTime"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	PicURLList      []string `json:"picURLList"`
	ThumpUp         int      `json:"thumpUp"`
	CommentListSize int      `json:"commendListSize"`
	CommentList     []retComment `json:"commendList"`
}

func (con TopicController) GetTopicDetail(context *gin.Context)  {
	//FIXME 判空 日期图片处理
	topicId := context.Param("topicId")
	var topic model.Topic
	result := core.DB.First(&topic, topicId)
	if result.Error != nil {
		log.Error(result.Error.Error())
		context.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
		return
	}
	var retTopicDetail topicDetail
	retTopicDetail.Status = "success"
	//发帖人相关信息
	var user model.User
	result = core.DB.First(&user, topic.Creator)
	if result.Error != nil {
		log.Error(result.Error.Error())
		context.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
		return
	}
	retTopicDetail.Nickname = user.Nickname
	//获得头像URL
	var portraitPic model.Picture
	result = core.DB.First(&portraitPic, user.PortraitId)
	if result.Error != nil {
		log.Error(result.Error.Error())
		context.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
		return
	}
	retTopicDetail.PortraitURL = portraitPic.PicURL

	retTopicDetail.CreateTime = topic.CreateAt.Format("2006-01-02")
	retTopicDetail.Title = topic.Title
	retTopicDetail.Content = topic.Content
	retTopicDetail.ThumpUp = topic.ThumpUp

	//获取返回的图片列表
	topicPicIds := strings.Split(topic.PicId, "|")
	picIds := make([]int, len(topicPicIds))
	for index, str := range topicPicIds{
		picIds[index],_ = strconv.Atoi(str)
	}
	var topicPics []model.Picture
	core.DB.Find(&topicPics, picIds)
	retPictureURLs := make([]string, len(topicPics))
	for index, pic := range topicPics{
		retPictureURLs[index] = pic.PicURL
	}
	retTopicDetail.PicURLList = retPictureURLs


	//返回评论相关
	var comments []model.Comment
	core.DB.Where(map[string]interface{}{"topic_id": topicId}).Find(&comments)
	var retComments = make([]retComment,len(comments))
	for index , comment := range comments{
		retComments[index].Content = comment.Content
		var commentUser model.User
		core.DB.First(&commentUser, comment.CommentusrId)
		retComments[index].Nickname = commentUser.Nickname
		var portraitPic model.Picture
		result = core.DB.First(&portraitPic, commentUser.PortraitId)
		if result.Error != nil {
			log.Error(result.Error.Error())
			context.JSON(http.StatusOK, gin.H{
				"status": "fail",
			})
			return
		}
		retComments[index].PortraitURL = portraitPic.PicURL
		retComments[index].ThumpUp = 12323
		retComments[index].CreateTime = "2021-21-12"
	}
	retTopicDetail.CommentList = retComments
	retTopicDetail.CommentListSize = len(retComments)

	context.JSON(http.StatusOK, retTopicDetail)
}

func (con TopicController) ReleaseTopic(context *gin.Context)  {

}
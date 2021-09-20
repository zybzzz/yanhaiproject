package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"yanhaiproject/core"
	"yanhaiproject/model"
	"yanhaiproject/tool"
)

type SchoolGroupController struct {
}

type retTopicInDetail struct {
	Title        string `json:"title"`
	Nickname     string `json:"nickname"`
	Portrait     string `json:"portrait"`
	CreateTime   string `json:"createTime"`
	Summary      string `json:"summary"`
	Tag          string `json:"tag"`
	ThumpUp      int `json:"thumpUp"`
	RecommendNum int `json:"recommendNum"`
}

type groupDetail struct {
	Status	string `json:"status"`
	SchoolPortraitURL string   `json:"schoolPortraitURL"`
	School            string   `json:"school"`
	IsAttention       bool   `json:"isAttention"`
	ResourceNum       int   `json:"resourceNum"`
	ResourceList      []string `json:"resourceList"`
	TopicNum          string   `json:"topicNum"`
	TopicList         []retTopicInDetail `json:"topicList"`
}

func (con SchoolGroupController) GetGroupList(context *gin.Context)  {

}


//FIXME  等待测试
func (con SchoolGroupController) GetGroupDetail(context *gin.Context)  {
	groupId := context.Param("groupId")
	userId := context.Query("context")

	//FIXME 等待上传接口与图片搭建完毕后修改
	var retGroupDetail groupDetail
	retGroupDetail.SchoolPortraitURL = "http://www.xxx.com"
	//查询出该小组相关信息
	var group model.Group
	iGroupId, _ := strconv.Atoi(groupId)
	core.DB.First(&group, iGroupId)
	retGroupDetail.School = group.GroupSchool
	retGroupDetail.ResourceNum = 2
	retGroupDetail.ResourceList = []string{retGroupDetail.School + "上岸技巧", retGroupDetail.School + "考研资料"}
	//查询用户是否关注
	var user model.User
	core.DB.First(&user, userId)
	if tool.IsContain(strings.Split(user.AttentionId,"|"), groupId) {
		retGroupDetail.IsAttention = true
	}else {
		retGroupDetail.IsAttention = false
	}
	//返回组中的帖子
	var topicsInGroup []model.Topic
	core.DB.Where(map[string]interface{}{"group_id":iGroupId}).Find(&topicsInGroup)
	retTopicInDetails := make([]retTopicInDetail, len(topicsInGroup))
	for index,topic := range topicsInGroup{
		retTopicInDetails[index].CreateTime = topic.CreateAt.Format("2006-01-02")
		retTopicInDetails[index].ThumpUp = topic.ThumpUp
		var creator model.User
		core.DB.First(&creator)
		retTopicInDetails[index].Nickname = creator.Nickname
		retTopicInDetails[index].Portrait = "http://xxx.xxx.com"
		retTopicInDetails[index].Title = topic.Title
		retTopicInDetails[index].Tag = topic.Tag
		//插入评论查询
		retTopicInDetails[index].RecommendNum = 1
		retTopicInDetails[index].Summary = topic.Content
	}
	retGroupDetail.TopicList = retTopicInDetails
	retGroupDetail.Status = "success"
	context.JSON(http.StatusOK,retGroupDetail)
}
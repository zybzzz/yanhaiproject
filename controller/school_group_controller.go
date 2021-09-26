package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"yanhaiproject/core"
	"yanhaiproject/model"
	"yanhaiproject/service"
	"yanhaiproject/tool"
	log "github.com/sirupsen/logrus"
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

type groupSummary struct {
	PortraitURL  string   `json:"portraitURL"`
	School       string   `json:"school"`
	ResourceNum  string   `json:"resourceNum"`
	ResourceList []string `json:"resourceList"`
	TopicNum     string   `json:"topicNum"`
	TopicList    []string `json:"topicList"`
}

type groupList struct {
	Status     string `json:"status"`
	HotListNum int `json:"hotListNum"`
	HotList    []groupSummary `json:"hotList"`
	GroupListNum int `json:"groupListNum"`
	GroupList    []groupSummary `json:"groupList"`
}
func (con SchoolGroupController) GetGroupList(context *gin.Context)  {
	//var list groupList
	//list.Status = "success"
	//list.HotListNum = 2
	//list.HotList = make([]groupSummary, list.HotListNum)
	////直接固定热度数据
	//list.HotList[0].School = "清华大学"
	//list.HotList[0].PortraitURL =

}


func (con SchoolGroupController) GetGroupDetail(context *gin.Context)  {
	groupId := context.Param("groupId")
	userId := context.Query("userId")

	var retGroupDetail groupDetail
	//retGroupDetail.SchoolPortraitURL = "http://www.xxx.com"
	//查询出该小组相关信息
	var group model.Group
	iGroupId, _ := strconv.Atoi(groupId)
	core.DB.First(&group, iGroupId)
	log.Info("查到的group是。。。")
	log.Info(group)
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
		//TODO 用户头像 等待测试
		retTopicInDetails[index].Portrait = service.PictureService{}.PicIdToURL(creator.PortraitId)
		retTopicInDetails[index].Title = topic.Title
		retTopicInDetails[index].Tag = topic.Tag
		//插入评论查询
		retTopicInDetails[index].RecommendNum = 1
		retTopicInDetails[index].Summary = topic.Content
	}
	retGroupDetail.TopicList = retTopicInDetails
	retGroupDetail.SchoolPortraitURL = service.PictureService{}.PicIdToURL(group.GroupPicId)
	retGroupDetail.Status = "success"
	context.JSON(http.StatusOK,retGroupDetail)
}
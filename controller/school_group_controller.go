package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"yanhaiproject/core"
	"yanhaiproject/model"
	"yanhaiproject/service"
	"yanhaiproject/tool"
)

type SchoolGroupController struct {
}

type retTopicInDetail struct {
	TopicId       int      `json:"topicId"`
	Title         string   `json:"title"`
	Nickname      string   `json:"nickname"`
	Portrait      string   `json:"portrait"`
	CreateTime    string   `json:"createTime"`
	Summary       string   `json:"summary"`
	Tag           []string   `json:"tag"`
	TopicPictures []string `json:"topicPictures"`
	ThumpUp       int      `json:"thumpUp"`
	RecommendNum  int      `json:"recommendNum"`
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
	GroupId      int      `json:"groupId"`
	PortraitURL  string   `json:"portraitURL"`
	School       string   `json:"school"`
	ResourceNum  int      `json:"resourceNum"`
	ResourceList []string `json:"resourceList"`
	TopicNum     int      `json:"topicNum"`
	TopicList    []string `json:"topicList"`
}

type groupList struct {
	Status     string `json:"status"`
	HotListNum int `json:"hotListNum"`
	HotList    []groupSummary `json:"hotList"`
	GroupListNum int `json:"groupListNum"`
	GroupList    []groupSummary `json:"groupList"`
}


//TODO 切换关注状态 等待测试
func (con SchoolGroupController) ChangeAttendStatus(context *gin.Context)  {
	userId := context.Param("userId")
	groupId := context.Param("groupId")
	var user model.User
	core.DB.First(&user, userId)
	userAttends := strings.Split(user.AttentionId, "|")
	var newAttends []string
	isExist := false
	for _, attendId := range userAttends{
		if attendId == groupId {
			isExist = true
		}else {
			newAttends = append(newAttends, attendId)
		}
	}

	if !isExist {
		newAttends = append(newAttends, groupId)
	}

	user.AttentionId = strings.Join(newAttends, "|")
	core.DB.Save(&user)

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}


func (con SchoolGroupController) GetGroupList(context *gin.Context)  {
	var list groupList
	list.Status = "success"
	list.HotListNum = 2
	list.HotList = make([]groupSummary, list.HotListNum)
	//直接固定热度数据
	//清华
	var hotgroup model.Group
	var hottopic model.Topic
	core.DB.First(&hotgroup, 1)
	list.HotList[0].GroupId = hotgroup.GroupId
	list.HotList[0].School = hotgroup.GroupSchool
	list.HotList[0].PortraitURL = service.PictureService{}.PicIdToURL(hotgroup.GroupPicId)
	list.HotList[0].ResourceNum = 1
	list.HotList[0].ResourceList = []string{"清华大学考研资料"}
	list.HotList[0].TopicNum = 1
	core.DB.Where(map[string]interface{}{"group_id": 1}).First(&hottopic)
	list.HotList[0].TopicList = []string{hottopic.Title}
	//北大
	var hotgroup2 model.Group
	var hottopic2 model.Topic
	core.DB.First(&hotgroup2, 2)
	list.HotList[0].GroupId = hotgroup2.GroupId
	list.HotList[1].School = hotgroup2.GroupSchool
	list.HotList[1].PortraitURL = service.PictureService{}.PicIdToURL(hotgroup2.GroupPicId)
	list.HotList[1].ResourceNum = 1
	list.HotList[1].ResourceList = []string{"北京大学考研资料"}
	list.HotList[1].TopicNum = 1
	core.DB.Where(map[string]interface{}{"group_id": 2}).First(&hottopic2)
	list.HotList[1].TopicList = []string{hottopic2.Title}

	//返回普通小组
	var groups []model.Group
	core.DB.Find(&groups)
	list.GroupListNum = len(groups)
	list.GroupList = make([]groupSummary, list.GroupListNum)
	for index,group := range groups{
		list.GroupList[index].GroupId = group.GroupId
		list.GroupList[index].School = group.GroupSchool
		list.GroupList[index].PortraitURL = service.PictureService{}.PicIdToURL(group.GroupPicId)
		list.GroupList[index].ResourceNum = 1
		list.GroupList[index].ResourceList = []string{group.GroupSchool + "考研资料"}
		list.GroupList[index].TopicNum = 1
		var topic model.Topic
		core.DB.Where(map[string]interface{}{"group_id": group.GroupId}).First(&topic)
		list.GroupList[index].TopicList = []string{topic.Title}
	}


	context.JSON(http.StatusOK, list)
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
		retTopicInDetails[index].TopicId = topic.TopicId
		retTopicInDetails[index].TopicPictures = service.PictureService{}.PicIdsToURL(topic.PicId)
		retTopicInDetails[index].CreateTime = topic.CreateAt.Format("2006-01-02")
		retTopicInDetails[index].ThumpUp = topic.ThumpUp
		var creator model.User
		core.DB.First(&creator)
		retTopicInDetails[index].Nickname = creator.Nickname
		//TODO 用户头像 等待测试
		retTopicInDetails[index].Portrait = service.PictureService{}.PicIdToURL(creator.PortraitId)
		retTopicInDetails[index].Title = topic.Title
		retTopicInDetails[index].Tag = strings.Split(topic.Tag,"|")
		//插入评论查询
		retTopicInDetails[index].RecommendNum = 1
		retTopicInDetails[index].Summary = topic.Content
	}
	retGroupDetail.TopicList = retTopicInDetails
	retGroupDetail.SchoolPortraitURL = service.PictureService{}.PicIdToURL(group.GroupPicId)
	retGroupDetail.Status = "success"
	context.JSON(http.StatusOK,retGroupDetail)
}
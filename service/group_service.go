package service

import (
	"strconv"
	"strings"
	"yanhaiproject/core"
	"yanhaiproject/model"
)

type GroupService struct {
}

/**
 * @author ZhangYiBo
 * @description  从一堆关注id中获取关注列表
 * @date 11:47 上午 2021/9/16
 * @param groupIds 传入的|字符串
 * @return 关注类型结构列表
 **/
func (GroupService) getAttentionList(groupIds string) (list []attention) {
	//TODO 接口完成，append方法待优化 待测试
	//list = append(list, attention{
	//	School:            "hahaha",
	//	SchoolPortraitURL: "123",
	//})
	allGroupIds := strings.Split(groupIds, "|")
	iGroupIds := make([]int, len(allGroupIds))
	for index, str := range allGroupIds{
		iGroupIds[index],_ = strconv.Atoi(str)
	}
	var groups []model.Group
	core.DB.Find(&groups, iGroupIds)
	for _, group := range groups{
		list = append(list, attention{
			GroupId:           group.GroupId,
			School:            group.GroupSchool,
			SchoolPortraitURL: PictureService{}.PicIdToURL(group.GroupPicId),
		})
	}
	return
}


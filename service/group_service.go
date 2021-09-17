package service

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
	//FIXME 接口待完成，append方法待优化
	list = append(list, attention{
		School:            "hahaha",
		SchoolPortraitURL: "123",
	})
	return
}


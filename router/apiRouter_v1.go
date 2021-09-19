/*
	包router单独抽离封装路由
 */
package router

import (
	"github.com/gin-gonic/gin"
	"yanhaiproject/controller"
)

//	路由前缀
const PREFIX_URL = "/api/v1"

func ApiV1RouterInit(engine *gin.Engine)  {
	apiV1Router := engine.Group(PREFIX_URL)
	{
		//	TODO 仍然缺省控制器实现

		//	获取推荐了列表
		apiV1Router.GET("/getRecommendList", controller.IndexPageController{}.GetRecommendList)
		//	获取关注列表
		apiV1Router.GET("/getAttentionList", controller.IndexPageController{}.GetAttentionList)
		//	获取圈子列表
		apiV1Router.GET("/getGroupList", controller.SchoolGroupController{}.GetGroupList)
		//	获取具体圈子信息
		apiV1Router.GET("/getGroupDetail/:groupId", controller.SchoolGroupController{}.GetGroupDetail)
		//	获取我的信息
		apiV1Router.GET("/getMyMessage", controller.UserController{}.GetMyMessage)
		//	更改我的信息
		apiV1Router.PUT("/changeMyMess", controller.UserController{}.ChangeMyMess)
		//	获取我发布的帖子
		apiV1Router.GET("/getMyReleaseList", controller.UserController{}.GetMyReleaseList)
		//	登录
		apiV1Router.GET("/login", controller.LoginController{}.Login)
		//	注册
		apiV1Router.POST("/register", controller.LoginController{}.Register)
		//	获取帖子详细信息
		apiV1Router.GET("/getTopicDetail/:topicId", controller.TopicController{}.GetTopicDetail)
		//	发布帖子
		apiV1Router.POST("/releaseTopic", controller.TopicController{}.ReleaseTopic)
		//	上传图片
		apiV1Router.POST("/uploadPicture", controller.PictureController{}.UploadPicture)
		//	评论
		apiV1Router.POST("/comment", controller.CommentController{}.Comment)
	}
}
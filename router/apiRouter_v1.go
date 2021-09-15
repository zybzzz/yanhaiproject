/*
	包router单独抽离封装路由
 */
package router

import "github.com/gin-gonic/gin"

//	路由前缀
const PREFIX_URL = "/api/v1"

func ApiV1RouterInit(engine *gin.Engine)  {
	apiV1Router := engine.Group(PREFIX_URL)
	{
		//	TODO 仍然缺省控制器实现

		//	获取推荐了列表
		apiV1Router.GET("/getRecommendList")
		//	获取关注列表
		apiV1Router.GET("/getAttentionList")
		//	获取圈子列表
		apiV1Router.GET("/getGroupList")
		//	获取具体圈子信息
		apiV1Router.GET("/getGroupDetail/:groupId")
		//	获取我的信息
		apiV1Router.GET("/getMyMessage")
		//	更改我的信息
		apiV1Router.PUT("/changeMyMess")
		//	获取我发布的帖子
		apiV1Router.GET("/getMyReleaseList/:userId")
		//	登录
		apiV1Router.GET("/login")
		//	注册
		apiV1Router.POST("/register")
		//	获取帖子详细信息
		apiV1Router.GET("/getTopicDetail/:topicId")
		//	发布帖子
		apiV1Router.GET("/releaseTopic")
		//	上传图片
		apiV1Router.POST("/uploadPicture")
		//	评论
		apiV1Router.POST("/recommend")
	}
}
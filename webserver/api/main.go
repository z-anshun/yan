package main

import (
	"github.com/gin-gonic/gin"
	"session"
)

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		if !validateUserSession(c) {
			c.Redirect(307, "/user")
		} //判断是否留有session
	}
}

func RegisterHandlers() *gin.Engine {
	e := gin.Default()
	e.Use(middleware())
	e.POST("/user", CreateUser)            //用户创建
	e.POST("/user/:user_name", Login)      //登录
	e.GET("/user/:user_name", GetUserInfo) //获取用户信息
	videoGroup := e.Group("/user/:user_name/videos")
	{
		videoGroup.POST("", AddNewVideo)           //视频上传
		videoGroup.GET("", ListAllVideos)          //获取视频信息
		videoGroup.DELETE("/:vid_id", DeleteVideo) //删除视频 只能删除自己滴
	}
	e.POST("/user/:user_name/videos/:vid-id/comments", PostComment) //发评论
	e.GET("/videos/:vid-id/comments", ShowComments)                 //展示评论
	e.POST("/user/:user_name/videos/:vid-id", AddVideoAchieve)      //用户点赞
	e.GET("videos/:vid-id", ShowVideoAchieve)                       //得到视频的
	e.DELETE("/user/:user_name/videos/", DeleteVideoAchieve)        //取消点赞等
	e.GET("/user/:user_name/videos/:vid-id", GetUserAchieveSelf)    //获取用户对于该视频是否点赞
	e.GET("/author/:name", GetUserAchieve)                          //获取用户的点赞数，粉丝，播放量等

	Fangoup := e.Group("/user/:user_name")
	{
		Fangoup.POST("/fan/:author", AddFan)
		Fangoup.DELETE("/fan", DelFan)
		Fangoup.GET("/fan", GetUserAtten)
	}
	return e
}
func Prepare() {
	session.LoadSessionFromDB()
}
func main() {
	Prepare() //那取所有的sessions
	r := RegisterHandlers()
	r.Run(":8000")
}

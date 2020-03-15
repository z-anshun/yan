package api

import (
	"github.com/gin-gonic/gin"
	"src/second_work/middle"
)



func SetRouter(e *gin.Engine) {
	//注册
	e.POST("/register", Register)
	//获取排行榜
	e.GET("/chart", Chart)
	e.Use(middle.Middleware())
	//登录
	e.POST("/login", Login)
	//参赛
	e.POST("/join", JoinMatch)
	//退赛
	e.DELETE("/leave", LeaveMatch)
	//投票
	e.POST("/vote", Vote)
	//下面是在10s中内
	g := e.Group("/final", middle.CheckTime()) //这个管道是全局的，所以问题不大
	{
		g.POST("/fvote", FinalVote)
		g.GET("/fresh",FreshChart)
		g.GET("/fchart", FinalChart)
	}

}

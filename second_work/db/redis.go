package db

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	. "src/second_work"
	"src/second_work/model"
	"strings"

	"src/second_work/limite"
	"src/second_work/resp"
	"strconv"
	"time"
)

func RedisSetIn(c *gin.Context, u model.User) {
	cmd, err := R.Exists(u.Name).Result()
	if err != nil {
		panic(err)
	}
	if cmd != 0 { //用户名重复
		resp.UserExist(c)
	}

	//转化为map
	mess := map[string]interface{}{
		"name":         u.Name,
		"password":     u.Password,
		"registertime": time.Now().Format("2006-01-02 15:04:05"), //注册时间
		"logintime":    time.Now().Format("2006-01-02 15:04:05"), //登录时间
		"membervote":   "3",                                      //投票数
		"final":        "false",                                  //最后是否投票，最后每人只能投一次
	}

	_, err = R.HMSet(u.Name, mess).Result()
	if err != nil {
		resp.RedisError(c, "redis set in error")
	}

}

//value获取
func UserGet(n string, k string) string {
	return R.HGet(n, k).Val()

}
func changeTime(s string)string {
	arr:=strings.Split(s,"-")
	str:=""
	for _,v:=range arr{
		str=str+v
	}
	return str

}
//更新票数
func upateVote(n string) {
	t := R.HGet(n, "logintime").Val()[0:10]
	now := time.Now().Format("2006-01-02 15:04:05")[0:10]
	t_i, err := strconv.Atoi(changeTime(t))
	if err != nil {
		log.Fatal("time to int error:", err)
	}
	now_i, err := strconv.Atoi(changeTime(now) ) //这里直接将2020-03-12 ——> 2020-03-12
	if err != nil {
		log.Fatal("time to int error:", err)
	}
	//时间比现在小一天多了
	if t_i < now_i {
		_, err := R.HSet(n, "membervote", "3").Result()
		if err != nil {
			log.Fatal("redis update vote error:", err)
		}
	}
	return
}

//登录时间更改
func LoginTimeUpdate(n string) {
	//更改登录时间前，先把投票数改了
	//因为redis无法便利哈希表的key，所有每次在登录后看是否要更新票数
	upateVote(n)

	_, err := R.HSet(n, "logintime", time.Now().Format("2006-01-02 15:04:05")).Result()
	if err != nil {
		log.Fatal("redis update error:", err)
	}

}

//加入redis
func JoinMat(c *gin.Context, n string) {
	count,_:=R.ZScore("compertitor",n).Result()
	if count!=0{
		resp.UserExist(c)
	}
	_, err := R.Do("ZADD", "competitor", 0, n).Result() //加入比赛
	if err != nil {
		resp.RedisError(c, "redis join match error")
	}

}

//删除redis
//错误干脆就直接处理了
func DeleteMat(c *gin.Context, n string) {
	_, err := R.HDel(n, "vote").Result()
	if err != nil {
		resp.RedisError(c, "redis leave match error")
	}
}

//投票
func PutVote(c *gin.Context, n string, obj string) {
	member_vote := R.HGet(n, "membervote").Val() //用户的信息里面的投票数
	m, err := strconv.Atoi(member_vote)
	if err != nil {
		log.Fatal("member to int error:", err)
	}
	if m <= 0 {
		c.JSON(200, gin.H{
			"code":    "007",
			"message": "you can vote any other",
		})
	}

	_, err = R.HSet(n, member_vote, strconv.Itoa(m-1)).Result()
	if err != nil {
		resp.RedisError(c, "reduces user vote error")
	}
	vote := R.ZScore("competitor", obj).Val()                 //先得到他的投票数
	_, err = R.Do("ZADD", "competitor", vote+1, obj).Result() //更改
	if err != nil {
		resp.RedisError(c, "add vote error")
	}

}

//get 排行榜
func GetChart(c *gin.Context) *[]redis.Z {
	//cou:=R.ZCard("competitor").Val()
	//获取所有，因为是集合，则直接
	com, err := R.ZRevRangeWithScores("competitor", 0, -1).Result()
	if err != nil {
		resp.RedisError(c, "get competitor chart error")
	}

	return &com
}

//最后的开启
func FinalMatch() {
	com, err := R.ZRange("competitor", 0, -1) .Result()
	if err != nil {
		log.Panic("creat final redis error")
	}
	for _, v := range com {
		_, err := R.Do("ZADD", "final", 0, v).Result()
		if err != nil {
			log.Fatal("creat final redis error")
		}
	}
}

//最后的投票
func FinalVoteUser(c *gin.Context, n string, obj string,coon *limite.ConnLimiter) {
	flag := R.HGet(n, "final").Val()
	if flag == "true" {
		resp.VotedUser(c)
	}else {
		err:=R.HSet(n, "final", "true").Err()
		if err!=nil{
			resp.RedisError(c,"update user final error")
		}
		vote:=R.ZScore("final",obj).Val()
		err=R.Do("ZADD","final",obj,vote+1).Err()
		if err!=nil{
			resp.RedisError(c,"add competitor vote  error")
		}
		//这里投完了，，才敢释放
		coon.ReleaseCoon()
	}
}
//刷新最后的统计
func FreshFinalChart(c *gin.Context)[]string{
	str,cursor,err:=R.ZScan("final",0,"*",100).Result()
	if err!=nil{
		resp.RedisError(c,"zscan redis error")
	}
	//如果没刷新完
	if cursor!=0{
		FreshFinalChart(c)
	}
	return  str

}
//获取最后的
func GetFinalChart(c *gin.Context)*[]redis.Z{
	users,err:=R.ZRangeWithScores("final",0,-1).Result()
	if err!=nil{
		resp.RedisError(c,"get final chart error")
	}
	return  &users
}
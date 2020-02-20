package main

import (
	"dbops"
	"defs"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"session"
	"strconv"
	"utils"
)

func CreateUser(c *gin.Context) {

	ubody := &defs.UserCredential{}
	ubody.Username = c.PostForm("user_name")
	ubody.Pwd = c.PostForm("pwd")

	//设置cookie方便前端获取
		c.SetCookie("user_name",ubody.Username ,60*60*24,"/","localhost",false,true )


	if len(ubody.Username) == 0 || len(ubody.Pwd) == 0 {
		sendErrorResponse(c.Writer, defs.ErrorRequestBodyParseFailed) //读取失败
		return
	}
	if err := dbops.AddUserCredential(ubody.Username,ubody.Pwd); err != nil {
		sendErrorResponse(c.Writer, defs.ErrorDBError) //添加数据库失败
		return
	}

	id := session.GenerateNewSessionId(ubody.Username) //产生一个session
	su := &defs.SignedUp{Success: true, SessionId: id} //写入session
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(c.Writer, defs.ErrorInternalFaults) //添加失败
		return
	} else {
		sendNormalResponse(c.Writer, string(resp), 201) //添加成功
	}

}
func Login(c *gin.Context) {


	ubody := &defs.UserCredential{}
	ubody.Username = c.PostForm("user_name")
	ubody.Pwd = c.PostForm("pwd")
	_, err:= c.Cookie("user_name")
	if err!=nil{
		//没有了，，就再设置
		c.SetCookie("user_name",ubody.Username ,60*60*24,"/","localhost",false,true )
	}
	if len(ubody.Username) == 0 || len(ubody.Pwd) == 0 {
		sendErrorResponse(c.Writer, defs.ErrorRequestBodyParseFailed) //读取失败
		return
	}
	//uname := c.PostForm("login_name")
	//if uname != ubody.Username {
	//	sendErrorResponse(c.Writer, defs.ErrorNotAuthUser)
	//	return
	//}
	pwd, err := dbops.GetUserCredential(ubody.Username) //验证是否合法
	if err != nil || len(pwd) == 0 || pwd != ubody.Pwd {
		sendErrorResponse(c.Writer, defs.ErrorNotAuthUser)
		return
	}
	if pwd=="0"{

	}
	id := session.GenerateNewSessionId(ubody.Username)
	si := &defs.SignedIn{true, id}
	if resp, err := json.Marshal(si); err != nil {
		sendErrorResponse(c.Writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(c.Writer, string(resp), 200)
	}
}

func GetUserInfo(c *gin.Context) {
	//if !ValidateUser(c.Writer, c.Request.Response) {
	//	log.Println("Unathorized user")
	//	return
	//}

	uname := c.Param("user_name")
	id, err := dbops.GetUser(uname) //得到user的id
	if err != nil || id == -1 {
		log.Println("Error in GetUserInfo:", err)
		return
	}
	ui := &defs.UserInfo{id}
	if resp, err := json.Marshal(ui); err != nil {
		sendErrorResponse(c.Writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(c.Writer, string(resp), 200)
	}
}
func AddNewVideo(c *gin.Context) {
	nvbody := &defs.NewVideo{}

	id, err:= strconv.Atoi(c.PostForm("author_id"))
	nvbody.AuthorId=id
	nvbody.Name   = c.PostForm("name")
	if err!=nil|| len(nvbody.Name) == 0{
		sendErrorResponse(c.Writer, defs.ErrorRequestBodyParseFailed) //读取失败
		return
	}

	vi, err := dbops.AddNewVideo(nvbody.AuthorId, nvbody.Name)

	if err != nil {
		log.Println("Error in AddNewVideo:", err)
		sendErrorResponse(c.Writer, defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(vi); err != nil {
		sendErrorResponse(c.Writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(c.Writer, string(resp), 201) //上传成功
	}
}

//查找所有得videos
func ListAllVideos(c *gin.Context) {


	uname := c.Param("user_name")
	vs, err := dbops.ListVideosInfo(uname, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Println("Error in ListAllVideos:", err)
		sendErrorResponse(c.Writer, defs.ErrorInternalFaults)
		return
	}

	vsi := &defs.VideosInfo{vs}
	if resp, err := json.Marshal(vsi); err != nil {
		sendErrorResponse(c.Writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(c.Writer, string(resp), 200)
	}

}

func DeleteVideo(c *gin.Context) {

	vid := c.Param("vid-id")
	username := c.Param("user_name")
	err := dbops.DeleteVideoInfo(vid, username)

	if err != nil {
		log.Println("Error in DeleteVideo:", err)
		sendErrorResponse(c.Writer, defs.ErrorDBError)
		return
	}
	sendNormalResponse(c.Writer, "", 204)
}

func PostComment(c *gin.Context) {

	reqBody, _ := ioutil.ReadAll(c.Request.Body)

	cbody := &defs.NewComment{}

	if err := json.Unmarshal(reqBody, cbody); err != nil {
		log.Println(err)
		sendErrorResponse(c.Writer, defs.ErrorRequestBodyParseFailed)
		return
	}

	vid := c.Param("vid-id")
	if err := dbops.AddNewComments(vid, cbody.AuthorId, cbody.Content, cbody.Vvid); err != nil {
		log.Println("Error in PostComment:", err)
		sendErrorResponse(c.Writer, defs.ErrorDBError)
	} else {
		sendNormalResponse(c.Writer, "ok", 201)
	}
}

func ShowComments(c *gin.Context) {


	vid := c.Param("vid-id")
	cm, err := dbops.ListCommens(vid, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Println("Error in ShowComments:", err)
		sendErrorResponse(c.Writer, defs.ErrorDBError)
		return
	}

	cms := &defs.Comments{cm}
	if resp, err := json.Marshal(cms); err != nil {
		sendErrorResponse(c.Writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(c.Writer, string(resp), 200)
	}
}

func AddVideoAchieve(c *gin.Context) {

	name := c.Param("user_name")
	kind := c.PostForm("kind")
	vid := c.Param("vid-id")
	err := dbops.AddAchieve(name, vid, kind)
	if err != nil {
		log.Println("Error in PostComment:", err)
		sendErrorResponse(c.Writer, defs.ErrorDBError)
	} else {
		sendNormalResponse(c.Writer, "ok", 201)
	}
}
func ShowVideoAchieve(c *gin.Context) {


	vid := c.Param("vid-id")
	vid_ac, err := dbops.GetVideoAch(vid)
	if err != nil {
		sendErrorResponse(c.Writer, defs.ErrorDBError)
	}
	resp, err := json.Marshal(vid_ac)
	if err != nil {
		sendErrorResponse(c.Writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(c.Writer, string(resp), 200)
	}

}

func DeleteVideoAchieve(c *gin.Context) {


	vid := c.PostForm("vid-id")
	name := c.Param("user_name")
	kind, _ := c.GetPostForm("kind")
	err := dbops.DeleteVideoAch(vid, name, kind)
	if err != nil {
		sendErrorResponse(c.Writer, defs.ErrorDBError)
	} else {
		sendNormalResponse(c.Writer, "", 204)
	}

}

func GetUserAchieveSelf(c *gin.Context) {

	vid := c.Param("vid")
	name := c.Param("name")
	user, err := dbops.GetUserSelf(name, vid)
	if err != nil {
		sendErrorResponse(c.Writer, defs.ErrorDBError)
	}
	resp, err := json.Marshal(user)
	if err != nil {
		sendErrorResponse(c.Writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(c.Writer, string(resp), 200)
	}
}

func GetUserAchieve(c *gin.Context) {

	name := c.Param("name")
	userAch, err := dbops.GetUserAch(name)
	if err != nil {
		sendErrorResponse(c.Writer, defs.ErrorDBError)
	}
	resp, err := json.Marshal(userAch)
	if err != nil {
		sendErrorResponse(c.Writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(c.Writer, string(resp), 200)
	}
}

func AddFan(c *gin.Context) {


	name := c.Param("user_name")
	author := c.Param("author")
	err := dbops.AddUserFan(name, author)
	if err != nil {
		sendErrorResponse(c.Writer, defs.ErrorDBError)
	} else {
		sendNormalResponse(c.Writer, "ok", 201)
	}
}

func DelFan(c *gin.Context) {


	name := c.Param("user_name")
	author, _ := c.GetQuery("author")
	err := dbops.DelUserFan(name, author)
	if err != nil {
		sendErrorResponse(c.Writer, defs.ErrorDBError)
	} else {
		sendNormalResponse(c.Writer, "", 204)
	}
}

func GetUserAtten(c *gin.Context) {

	name := c.Param("user_name")
	atten, err := dbops.GetUserAtten(name)
	if err != nil {
		sendErrorResponse(c.Writer, defs.ErrorDBError)
	}
	attens := defs.Attens{atten}
	resp, err := json.Marshal(attens)
	if err != nil {
		sendErrorResponse(c.Writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(c.Writer, string(resp), 200)
	}

}

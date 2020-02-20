package main

import (
	"github.com/gin-gonic/gin"
	"dbops"
	"session"
)

//判断用户注册校验


//检查head 在不在
func validateUserSession(c *gin.Context) bool {
	name,_:= c.GetPostForm("user_name")
	if len(name) == 0 {
		return false
	}
	sess,err:=dbops.RetrieveSession(name)
	if err!=nil{
		return  false
	}
	_, ok := session.IsSessinExpired(sess.Sess)//就判断session id 呗
	if ok {
		return false
	}

	return true
}
//func ValidateUser(w http.ResponseWriter, r *http.Response) bool {
//	uname := r.Header.Get(HEADER_FIELD_UNAME)
//	if len(uname) == 0 {
//		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
//		return false
//	}
//	return true
//}

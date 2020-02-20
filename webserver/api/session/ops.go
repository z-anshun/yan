package session

import (
	"dbops"
	"defs"
	"sync"
	"time"
	"utils"
)

//sync.Map 实现线程安全，在读上面很nb 写上一般
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}
func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}
func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

//从数据库中拿数据
func LoadSessionFromDB() {

	r, err := dbops.RetrieveAllSessions() //拿到所有
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})

}

//用户登录时的缓存
func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()     //毫秒
	ttl := ct + 30*60*1000 //存储时间
	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)
	return id
}

//判断用户是否过期等
func IsSessinExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid) //得到缓存里面得那个用户
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}

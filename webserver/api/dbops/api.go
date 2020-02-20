package dbops

import (
	"database/sql"
	"defs"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"time"
	"utils"

)

//增加
func AddUserCredential(loginName string, pwd string) error {
	p,_:= GetUserCredential(loginName)
	if p!="" {
		return  errors.New("昵称重复")
	}
	stmtIN, err := dbConn.Prepare("insert into user (login_name,pwd) values(?,?) ")

	if err != nil {
		return err
	}
	//defer stmtIN.Close() 出栈时才会用，对性能有损耗
	_, err = stmtIN.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIN.Close()
	return nil
}

//查
func GetUserCredential(loginName string) (string, error) {
	stmtOUT, err := dbConn.Prepare("select pwd from user where login_name=?")
	if err != nil{
		log.Printf("%s", err)
		return "", err
	}
	//defer dbConn.Close()
	var pwd string
	err = stmtOUT.QueryRow(loginName).Scan(&pwd)
	if err==sql.ErrNoRows{
		return  "0",err
	}
	//sql.ErrNorows为没有结果
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOUT.Close()
	return pwd, nil
}

//删
func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from user where login_name=? and pwd=?")
	if err != nil {
		log.Println("DeleteUser error:", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

//get
func GetUser(name string) (int, error) {
	stmtGet, err := dbConn.Prepare("select id from user where login_name=?")
	if err != nil || err == sql.ErrNoRows {
		log.Println("GetUser error:", err)
		return -1, err
	}
	var id int
	err = stmtGet.QueryRow(name).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return -1, err
	}
	defer stmtGet.Close()
	return id, nil
}

//video add
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, _ := utils.NewUUID()
	//creatime->db 此处为display time
	//aid 用户id name名字 vid视频id
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:04") //M D Y,HH:MM:SS 上传的时间
	stmtIns, err := dbConn.Prepare(`insert into video_info 
(id,author_id,name,display_ctime) values(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime} //传递址
	defer stmtIns.Close()
	return res, nil
}

//video get
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("select author_id,name,display_ctime from video_info where id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return res, nil
}

//delete
func DeleteVideoInfo(vid string, username string) error {

	id, _ := GetUser(username)

	stmtDel, err := dbConn.Prepare("delete from video_info where id=?and author_id=?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid, id)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func ListVideosInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	id, err := GetUser(uname)
	if err != nil {
		return nil, err
	}
	stmtLists, err := dbConn.Prepare(`SELECT id,NAME,display_ctime FROM video_info
	WHERE author_id=1 AND  video_info.display_ctime>FROM_UNIXTIME(0) AND video_info.display_ctime<=FROM_UNIXTIME(0)
	ORDER BY video_info.display_ctime DESC`) //按时间排序

	var res []*defs.VideoInfo
	if err != nil {
		return nil, err
	}
	rows, err := stmtLists.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var vid, name, diaplay_ctime string
		err := rows.Scan(&vid, &id, &name, &diaplay_ctime)
		if err != nil {
			return res, err
		}
		v := &defs.VideoInfo{vid, id, name, diaplay_ctime}
		res = append(res, v)
	}
	defer stmtLists.Close()
	return res, nil
}

//comments add
func AddNewComments(vid string, aid int, content string, vvid int) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIN, err := dbConn.Prepare(`insert into comment
(id,video_id,author_id,content,vvid) values (?,?,?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stmtIN.Exec(id, vid, aid, content, vvid)
	if err != nil {
		return err
	}
	defer stmtIN.Close()
	return nil
}

//读 comments
func ListCommens(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`select comment.vid,comment.id,user.login_name,comment.content
from comment INNER JOIN user ON comment.author_id=user.id where comment.video_id=? and comment.vvid=0 and 
comment.time>FROM_UNIXTIME(?) and comment.time<=FROM_UNIXTIME(?)
ORDER BY comment.time DESC `) //按时间排序 这里拿出的是直接评论视频的
	var res []*defs.Comment
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var v_vid, id, name, content string
		if err := rows.Scan(&v_vid, &id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c) //把每个获得的评论压进去
		//查找评论的评论 v_id就是评论自增的那个id用它来查找vvid=vid
		v_id, _ := strconv.Atoi(v_vid)
		c.Next.Comments, err = ListReturnComments(vid, v_id, from, to)
		if err != nil {
			return res, err
		}
	}
	defer stmtOut.Close()
	return res, nil
}

//vid从1开始增加 vvid表示评论的回复的那个vid 默认为0
func ListReturnComments(vid string, v_vid int, from, to int) ([]*defs.Comment, error) {
	stmtOuts, err := dbConn.Prepare(`select comment.vid,comment.id,user.login_name,comment.content from comment
inner join user on comment.author_id=user.id where comment.video_id=? and comment.vvid=? and 
comment.time>FROM_UNIXTIME(?) and comment.time<=FROM_UNIXTIME(?) ORDER BY comment.time DESC
`) //读取回复这个评论的评论
	var res []*defs.Comment
	rows, err := stmtOuts.Query(vid, from, to, v_vid)
	for rows.Next() {
		var id, name, content string
		var v_id int
		if err := rows.Scan(&v_id, &id, &name, &content); err != nil {
			return res, err
		}
		c := defs.Comment{Id: id, VideoId: vid, Author: name, Content: content, Next: nil}
		res = append(res, &c) //把每个获得的评论压进去
		if v_id != 0 {
			c.Next.Comments, err = ListReturnComments(vid, v_id, from, to)
			//如果评论里面有回复这个评论的，就实现链表，链表里面链的是一陀评论
			if err != nil {
				return res, err
			}
		} else {
			return res, nil
		}
	}
	defer stmtOuts.Close()
	return res, nil
}

//kind只能是favour赞 collect收藏 和transmit转发
func AddAchieve(name string, vid string, kind string) error {
	kind = SwitchKind(kind) //转化为正则
	stmtAdd, err := dbConn.Prepare("update video_achieve set ?=? +1 where vid=?")
	if err != nil {
		return nil
	}
	_, err = stmtAdd.Exec(kind, kind, vid)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		stmtAdd, err = dbConn.Prepare("insert into video_achieve (?,vid) values(?,?)")
		_, err = stmtAdd.Exec(kind, 1, vid)
		if err != nil {
			return err
		}
	}
	err = addUserSelf(name, vid, kind)
	if err != nil {
		return err
	}
	defer stmtAdd.Close()
	return err
}

//name就是点赞或者干傻子的这个人
func addUserSelf(name string, vid string, kind string) error {
	kind = SwitchKind(kind) //转化为正则
	stmtAdd, err := dbConn.Prepare("update user-collect set ?=? +1 where vid=? and name=?")
	if err != nil {
		return err
	}
	_, err = stmtAdd.Exec(kind, kind, vid, name)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		stmtAdd, err = dbConn.Prepare("insert into video_achieve (?,vid,name) values(?,?,?)")
		_, err = stmtAdd.Exec(kind, 1, vid, name)
		if err != nil {
			return err
		}
	}
	defer stmtAdd.Close()
	return err
}

func GetVideoAch(vid string) (*defs.VideoAchieve, error) {
	stmtGet, err := dbConn.Prepare("seletc favour,transmit,collect from video_achieve where vid=?")
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	vid_ach := defs.VideoAchieve{vid, 0, 0, 0}
	if err == sql.ErrNoRows {
		return &vid_ach, nil
	}
	err = stmtGet.QueryRow(vid).Scan(&vid_ach.Favour, &vid_ach.Transmit, &vid_ach.Collect)
	if err != nil {
		return nil, err
	}
	defer stmtGet.Close()
	return &vid_ach, nil

}

//对视频 用户 还有自己都有影响
func DeleteVideoAch(vid string, name string, kind string) error {
	kind = SwitchKind(kind) //转化为正则

	stmtDel, err := dbConn.Prepare(`update video_achieve set ?=?-1 where vid=?;
update user-collect set ?=?-1 where vid=?;update user_achieve set ?=?-1 where name=? `)
	if err != nil { //这个删应该就没得找不出的可能
		return err
	}
	_, err = stmtDel.Exec(kind, kind, vid, kind, kind, vid, kind, kind, name)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func GetUserSelf(name string, vid string) (*defs.UserSelf, error) {
	stmeGet, err := dbConn.Prepare("select * from user-collect where name=? and vid=?")
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	user_collect := defs.UserSelf{name, vid, 0, 0, 0}
	if err == sql.ErrNoRows {
		return &user_collect, nil
	}
	err = stmeGet.QueryRow(name, vid).Scan(&user_collect.Name, &user_collect.Vid, &user_collect.Collect, &user_collect.Favour, &user_collect.Transmit)
	if err != nil {
		return nil, err
	}
	defer stmeGet.Close()
	return &user_collect, nil
}

func GetUserAch(name string) (*defs.UserAchieve, error) {
	stmeGet, err := dbConn.Prepare("select * from user_achieve where name=?")
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	userAch := defs.UserAchieve{name, 0, 0, 0}
	if err == sql.ErrNoRows {
		return &userAch, nil
	}
	err = stmeGet.QueryRow(name).Scan(&userAch.Name, &userAch.Favour, &userAch.Fan, &userAch.Transmit)
	if err != nil {
		return nil, err
	}
	defer stmeGet.Close()
	return &userAch, nil
}
func AddUserFan(name string, author string) error {
	stmtAdd, err := dbConn.Prepare(`insert into user_atten (name,atten) values(?,?);
update user_acieve set fan=fan+1 where name=?`)
	if err != nil {
		return err
	}
	_, err = stmtAdd.Exec(name, author, author)
	if err != nil {
		return nil
	}
	return nil
}

func DelUserFan(name string, author string) error {
	stmtDel, err := dbConn.Prepare(`delete   from user_atten where name=? and atten=?;
update user_acieve set fan=fan-1 where name=?`)
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(name, author, author)
	if err != nil {
		return nil
	}
	return nil
}

func GetUserAtten(name string) ([]*defs.Atten, error) {
	stmeGet, err := dbConn.Prepare(`select atten from user_atten where name=?`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	var authors []*defs.Atten
	if err == sql.ErrNoRows {
		return authors, err
	}
	rows, err := stmeGet.Query(name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a *defs.Atten
		err = rows.Scan(a.AuthorName)
		if err != nil {
			return authors, err
		}
		authors = append(authors, a)
	}
	defer stmeGet.Close()
	return authors, nil
}

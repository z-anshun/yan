package dbops

import (
	"database/sql"
	"defs"
	"log"
	"strconv"
	"sync"
)

//存放与db之间的交流 给session 拿session 删session

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("insert into sessions (session_id,TTL,login_name) values(?,?,?) ")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

//得一条
func RetrieveSession(login_name string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select TTL,session_id from sessions login_name=?")
	if err != nil {
		return nil, err
	}
	var ttl,sid string
	err = stmtOut.QueryRow(login_name).Scan(&ttl, &sid)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer stmtOut.Close()
	var ttlint int64
	//ttlint 10进制 64位
	if ttlint, err = strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = ttlint
		ss.Username =login_name
		ss.Sess=sid
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}

//得all
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("select * from sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	for rows.Next() {
		var id, ttlstr, login_name string
		er := rows.Scan(&id, &ttlstr, &login_name)
		if er != nil {
			log.Printf("retrive sessions error:%s", er)
			break
		}
		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 != nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
			//log.Printf("session id:%s, ttl:%d",id,ss.TTL )
		}
	}
	defer stmtOut.Close()
	return m, nil
}
func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("delete from sessions where session_id=?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(sid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

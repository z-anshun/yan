package dbops

import (

	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddVdeoDeletionRecord(name string ){
		if name[0:2]=="Av"{
			deleteVideobyId(name)

		}else{
			deleteVideobyName(name)

		}
}
func deleteVideobyId(vid string )error{
	stmtIns,err:=dbConn.Prepare("insert into video_del_rec (video_id) value(?)")
	if err!=nil{
		return  err
	}
	_,err=stmtIns.Exec(vid)
	if err!=nil{
		log.Println("AddVideoDeletionRecord error:",err)
		return  err
	}
	defer  stmtIns.Close()
	return  nil
}
func deleteVideobyName(name string )error{
	stmtIns,err:=dbConn.Prepare("select id from video_info where name=?")
	if err!=nil{
		return err
	}
	var vid string
	err = stmtIns.QueryRow(name).Scan(&vid)
	if err!=nil{
		return  err
	}
	stmtIns.Close()
	deleteVideobyId(vid)
}
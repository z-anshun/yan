package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)


func clearTables(){
	dbConn.Exec("truncate user")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comment")
	dbConn.Exec("truncate sessions")
}
func TestMain(m *testing.M )  {
	clearTables()//初始化
	m.Run()
	clearTables()
}
func TestUserWorkFlow(t *testing.T){
	t.Run("add",testAddUser )
	t.Run("get",testGetUse )
	t.Run("del",testDeleteUser )
	t.Run("rege",testRegetUser )
}
func testAddUser(t *testing.T) {
	err:=AddUserCredential("as","123")
	if err!=nil{
		t.Errorf(err.Error() )
	}
}
func testGetUse(t *testing.T) {
	pwd,err:=GetUserCredential("as")
	if err!=nil||pwd!="123"{
		t.Errorf(err.Error() )
	}
}
func testDeleteUser(t *testing.T) {
	err:=DeleteUser("as","123")
	if err!=nil{
		t.Errorf(err.Error() )
	}
}
func testRegetUser(t *testing.T){
	pwd,err:=GetUserCredential("as")
	if err!=nil{
		t.Errorf(err.Error() )
	}
	if pwd!=""{
		t.Errorf("deleting user is failth")
	}
}
var tempvid string
func TestVideoWorkFlow(t *testing.T){
	clearTables()
	t.Run("Pre",testAddUser )
	t.Run("add",testAddVideoInfo  )
	t.Run("get",testGetVideoInfo  )
	t.Run("del",testDeleteVideoInfo  )
	t.Run("rege",testRegetVideo  )
}
func testAddVideoInfo(t *testing.T){
	vi,err:=AddNewVideo(1,"my-video")
	if err!=nil{
		t.Errorf(err.Error()  )
	}
	tempvid=vi.Id
}
func testGetVideoInfo(t *testing.T){
	_,err:=GetVideoInfo(tempvid)
	if err!=nil{
		t.Errorf(err.Error() )
	}
}
func testDeleteVideoInfo(t *testing.T){
	err:=DeleteVideoInfo(tempvid)
	if err!=nil{
		t.Errorf(err.Error() )
	}
}
func testRegetVideo(t *testing.T){
	_,err:=GetVideoInfo(tempvid)
	if err!=nil{
		t.Errorf(err.Error()  )
	}
}
func TestComments(t *testing.T){
	clearTables()
	t.Run("Pre",testAddUser )
	t.Run("add",testAddCommens )
	t.Run("list",testListCommens )
}
func testAddCommens(t *testing.T){
	vid:="123"
	aid:=1
	content:="i am"
	err:=AddNewComments(vid,aid,content )
	if err!=nil{
		t.Errorf(err.Error() )
	}
}
func testListCommens(t *testing.T){
	vid:="123"
	from:=1534343435
	to,_:=strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/100000000,10) )
	res,err:=ListCommens(vid,from,to)
	if err!=nil{
		t.Errorf("list:%s",err.Error() )
	}
	for i,ele:=range res{
		fmt.Println(i,ele)
	}
}
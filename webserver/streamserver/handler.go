package main

import (
	"github.com/gin-gonic/gin"

	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func testPageHandler(c *gin.Context) {
	t, _ := template.ParseFiles("./videos/upload.html")
	//sendErrorResponse(c.Writer ,200,"test" )
	t.Execute(c.Writer, nil)
}

func streamHander(c *gin.Context) {
	vid := c.Param("vid-id")
	vl := VIDEO_DIR + vid     //得到文件
	video, err := os.Open(vl) //打开文件
	if err != nil {
		log.Printf("Error when try to open file ", err.Error())
		sendErrorResponse(c.Writer, http.StatusInternalServerError, "Internal Error")
		return
	}
	c.Header("Content-Type", "video/mp4")
	http.ServeContent(c.Writer, c.Request, "", time.Now(), video)
	defer video.Close()

}

func uploadHandler(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MAX_UPLOAD_SIZE) //单位为bits
	if err := c.Request.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {           //限制文件大小
		sendErrorResponse(c.Writer, http.StatusBadRequest, "File is to big")
		return
	}
	file, _, err := c.Request.FormFile("file") //<from name="file"
	if err != nil {
		sendErrorResponse(c.Writer, http.StatusInternalServerError, "Internal is error")
		return
	}
	data, err := ioutil.ReadAll(file) //读取
	if err != nil {
		log.Println("Read file error:", err.Error())
		sendErrorResponse(c.Writer, http.StatusInternalServerError, "Read file error")
		return
	}
	fn := c.Param("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666) //写入 文件名，数据，权限 2进制写入
	if err != nil {
		log.Println("Write file error:", err.Error())
		sendErrorResponse(c.Writer, http.StatusInternalServerError, "Write file error")
		return
	}
	c.Writer.WriteHeader(http.StatusCreated)
	c.Writer.WriteString("Upload is successful")
}
func uploadAv(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MAX_UPLOAD_SIZE) //单位为bits
	if err := c.Request.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {           //限制文件大小
		sendErrorResponse(c.Writer, http.StatusBadRequest, "File is to big")
		return
	}
	file, _, err := c.Request.FormFile("file_av") //<from name="file_av"
	if err != nil {
		sendErrorResponse(c.Writer, http.StatusInternalServerError, "Internal is error")
		return
	}
	data, err := ioutil.ReadAll(file) //读取
	if err != nil {
		log.Println("Read file error:", err.Error())
		sendErrorResponse(c.Writer, http.StatusInternalServerError, "Read file error")
		return
	}
	fn := c.Param("user_name")
	err = ioutil.WriteFile(AV_DIR+fn, data, 0666) //写入 文件名，数据，权限 2进制写入
	if err != nil {
		log.Println("Write file error:", err.Error())
		sendErrorResponse(c.Writer, http.StatusInternalServerError, "Write file error")
		return
	}
	c.Writer.WriteHeader(http.StatusCreated)
	c.Writer.WriteString("Upload is successful")
}
func getAv(c *gin.Context) {
	name := c.Param("user_name")
	av := AV_DIR+name    //得到文件
	img, err := os.Open(av) //打开文件
	if err != nil {
		log.Printf("Error when try to open file ", err.Error())
		sendErrorResponse(c.Writer, http.StatusInternalServerError, "Internal Error")
		return
	}
	c.Header("Content-Type", "img/png/jpg")
	http.ServeFile(c.Writer, c.Request,av)
	defer img.Close()

}
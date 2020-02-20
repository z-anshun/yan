package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//进行代理
var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, c *gin.Context) {
	var resp *http.Response
	var err error

	switch b.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", b.Url, nil) //转化
		req.Header = c.Request.Header
		resp, err = httpClient.Do(req) //发送request  让前端js中收到的response与后端完全一致
		if err != nil {
			log.Println(err)
			return
		}
		normalResponse(c.Writer, resp)
	case http.MethodPost:
		req, _ := http.NewRequest("POST", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
		req.Header = c.Request.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		normalResponse(c.Writer, resp)
	case http.MethodDelete:
		req, _ := http.NewRequest("Delete", b.Url, nil)
		req.Header = c.Request.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		normalResponse(c.Writer, resp)
	default:
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.WriteString("Bad api request")
		return
	}
}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
		return
	}
	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}

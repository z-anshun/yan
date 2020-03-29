package rrf

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"strings"
)

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		w:          w,
		req:        req,
		queryParam: parseQuery(req.RequestURI),
		formParam:  getForm(req),
	}
}

func parseQuery(uri string) (res map[string]string) {
	res = make(map[string]string)
	uris := strings.Split(uri, "?")
	if len(uris) == 1 {
		//后面没跟参数这些

		return
	}
	param := uris[len(uris)-1]
	params := strings.Split(param, "&")
	for _, v := range params {
		vPair := strings.Split(v, "=")
		if len(vPair) == 1 {
			//这个表示未获取到
			vPair[1] = ""
		}
		res[vPair[0]] = vPair[1]
	}

	return
}

func getForm(r *http.Request) *url.Values {
	//设置最大读取的
	//直接
	//究极万金油
	r.ParseMultipartForm(128)
	//获取form
	form := r.PostForm

	return &form
}

func (c *Context) BindJson(v interface{}) error {
	postform := c.formParam

	var form string

	for form, _ = range *postform {
		if len(form) != 0 {
			break
		}
	}

	if len(form) == 0 {
		parseform := c.req.Form
		for form, _ = range parseform {
			if len(form) != 0 {
				break
			}
		}
	}

	if err := json.Unmarshal([]byte(form), v); err != nil {
		log.Println("bind json error")
		return err
	}
	return nil
}

func (c *Context) PostFrom(key string) string {
	n := c.formParam

	return n.Get(key)
}
func (c *Context) Query(key string) string {
	v := c.queryParam

	return v[key]
}

func (c *Context) Write(content string) {
	_, _ = c.w.Write([]byte(content))
}

func (c *Context)Next(){

}
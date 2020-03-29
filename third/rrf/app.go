package rrf

import (
	"log"
	"net/http"
	"strings"
)

func Default() *App {
	return &App{
		router: make(map[string]handlerMap),
	}
}

//app "post"->uri->handle
//“post”->uri->middle->handle

func (a *App) GET(uri string, handlers ...Handler) {
	a.handle("GET", uri, handlers)
}

func (a *App) POST(uri string, handlers ...Handler) {
	a.handle("POST", uri, handlers)
}

func (a *App) Group(uri string) *App {
	//如果没有group在a里填一个group
	var group *App
	if len(a.group) == 0 {
		group = &App{
			router: make(map[string]handlerMap),
			group:  uri,
		}
	}
	//如果有了
	group = &App{
		router: make(map[string]handlerMap),
		group:  a.group + uri,
	}
	//group->a
	group.otherApp = a
	return group
}

func (a *App) Use(handler Handler) {
	a.middle = append(a.middle, handler)
}

func (a *App) handle(method, uri string, handlers []Handler) {
	//如果这个是指向真正运行的a的
	if a.otherApp != nil {
		if len(a.group) != 0 {
			uri = a.group + uri
		}
		a.otherApp.handle(method, uri, handlers)
	} else {//只有当不是group的时候，也就是没有指向运行的a时，可真正赋值
		//获取到相应的handler
		h, ok := a.router[method]
		//没得就给它一个
		//这里相当于有重复的post或者get这些
		if !ok {
			m := make(handlerMap)
			a.router[method] = m
			h = m
		}

		//获取相应得handler,判断是否make

		_, ok = h[uri]
		if ok {
			//route相同了
			panic("same route")
		}

		//每个get，每个post  ->  对着一个路由
		h[uri] = handlers
	}
}

//启动路由
func (a *App) Run(port string) {
	http.Handle("/", a)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println("start serve error:", err)
	}

}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	httpMethod := req.Method
	uri := req.RequestURI //uri就是端口后面那些
	uris := strings.Split(uri, "?")

	if len(uris) < 1 {
		//有问题的
		return
	}
	//map是个好东西
	handler, ok := a.router[httpMethod]
	if !ok {
		log.Panicln("may be a hacker:", req.RemoteAddr)
		return
	}

	h, ok := handler[uris[0]]
	if !ok {
		Handler404(w, req)
		return
	}
	//中间件应该要在这里开始
	if len(a.middle) != 0 {
		for _, v := range a.middle {
			m := NewContext(w, req)
			v(m)
		}
	}
	for _, v := range h {
		//创建一个新的context
		c := NewContext(w, req)
		//开始执行那些
		v(c)
	}
}

//404
func Handler404(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("404 not find"))
}

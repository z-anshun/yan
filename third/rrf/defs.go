package rrf

import (
	"net/http"
	"net/url"
)

type Context struct {
	req        *http.Request
	w          http.ResponseWriter
	queryParam map[string]string
	formParam  *url.Values
}

type Handler func(*Context)

type handlerMap map[string][]Handler

type App struct {
	router map[string]handlerMap
	group string
	middle []Handler
	otherApp *App
}

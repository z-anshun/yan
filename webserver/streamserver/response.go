package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter,sc int,errMag string   ){
	w.WriteHeader(sc)
	io.WriteString(w,errMag)
}

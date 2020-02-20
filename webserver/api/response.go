package main

import (
	"defs"
	"encoding/json"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSc)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}
func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w,resp )
}

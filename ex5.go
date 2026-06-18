package main

import (
	"fmt"
	"net/http"
)

func Agenthandler(w http.ResponseWriter, r *http.Request) {
	result:=r.Header.Get("User-Agent")
	fmt.Fprintf(w,result)
}

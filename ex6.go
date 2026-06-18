package main

import (
	"fmt"
	"net/http"
)

func Dashhandler(w http.ResponseWriter, r *http.Request) {

	api := r.Header.Get("X-API-Key")
	hapi := "secret123"

	if hapi != api {
		w.WriteHeader(401)
		return
	}else{
		fmt.Fprintln(w,"Welcome")
	}
}

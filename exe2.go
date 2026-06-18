package main

import (
	"fmt"
	"net/http"
)


func Hellohandler(w http.ResponseWriter, r *http.Request){
	if r.Method!="GET"{
		w.WriteHeader(405)
		return
	}
	name:="Guest"
	query:=r.URL.Query().Get("name")
	if query!=""{
		name=query
	}
	fmt.Fprintf(w, "Hello, %s!",name)
}


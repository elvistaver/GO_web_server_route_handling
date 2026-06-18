package main

import("net/http"
"fmt")


func Pinghandler(w http.ResponseWriter, r *http.Request){

	if r.Method=="GET"{
		w.WriteHeader(200)
		fmt.Fprintln(w, "pong")
	}
	if r.URL.Path !="/ping"{
		w.WriteHeader(404)
		fmt.Fprintln(w,`<h2>site unavailable`)
		return
	}
}

func main(){
	http.HandleFunc("/ping", Pinghandler)
	http.HandleFunc("/hello", Hellohandler)
	http.HandleFunc("/count", COunthandler)
	http.HandleFunc("/calculate", Calchandler)
	http.HandleFunc("/agent", Agenthandler)
	http.HandleFunc("/dashboard", Dashhandler)
	http.HandleFunc("/legacy", Legahanler)
	http.HandleFunc("/v2", Legahanler)

	fmt.Println("server running...http://localhost:8080/ping")
	if err:= http.ListenAndServe(":8080",nil); err!=nil{
		fmt.Println("error running server")
	}
}
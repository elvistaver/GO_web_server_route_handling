package main

import (
	"fmt"
	"io"
	"net/http"
)

func COunthandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fmt.Fprintln(w, "Send a POST request with text to count words")
	}
	if r.Method == "POST" {
		count, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		defer r.Body.Close()
		result := string(count)

		fmt.Fprintf(w, "count %v", len(result))
	}
}

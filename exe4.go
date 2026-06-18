package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Calchandler(w http.ResponseWriter, r *http.Request) {

	result := 0

	aa := r.URL.Query().Get("a")
	opp := r.URL.Query().Get("op")
	bb := r.URL.Query().Get("b")

	if opp == "add" {
		n1, err := strconv.Atoi(aa)
		if err != nil {
			w.WriteHeader(400)
		}
		n2, err := strconv.Atoi(bb)
		if err != nil {
			w.WriteHeader(400)
		}
		result = n1 + n2
	}
	if opp == "subtract" {
		n1, err := strconv.Atoi(aa)
		if err != nil {
			w.WriteHeader(400)
		}
		n2, err := strconv.Atoi(bb)
		if err != nil {
			w.WriteHeader(400)
		}
		result = n1 - n2
	}
	if opp == "multiply" {
		n1, err := strconv.Atoi(aa)
		if err != nil {
			w.WriteHeader(400)
		}
		n2, err := strconv.Atoi(bb)
		if err != nil {
			w.WriteHeader(400)
		}
		result = n1 * n2
	}

	fmt.Fprintf(w, "Result: %d", result)
}

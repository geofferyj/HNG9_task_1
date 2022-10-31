package main

import (
	"fmt"
	"net/http"
)

// return a json response
func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"slackUsername": "%s", "backend": %t, "age": %d, "bio": "%s"}`, "geofferyj", true, 27, "I am geofferyj, first of my name, the debugger of code, and all round bad guy.")
}

func main() {
	http.HandleFunc("/", getHandler)
	http.ListenAndServe(":80", nil)
}

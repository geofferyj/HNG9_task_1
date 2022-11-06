package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type DataIn struct {
	OperationType string `json:"operation_type"`
	X int `json:"x"`
	Y int `json:"y"`
}

type DataOut struct {
	Result int `json:"result"`
	SlackUsername string `json:"slackUsername"`
	OperaionType string `json:"operation_type"`
}

// return a json response
func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"slackUsername": "%s", "backend": %t, "age": %d, "bio": "%s"}`, "geofferyj", true, 27, "I am geofferyj, first of my name, the debugger of code, and all round bad guy.")
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}
// endpoint to receive and add numbers 
func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data DataIn

	json.NewDecoder(r.Body).Decode(&data)

	var result int
	var status string
	switch data.OperationType {
	case "addition":
		result, status = add(data.X, data.Y), "success"
	case "subtraction":
		result, status = subtract(data.X, data.Y), "success"
	case "multiplication":
		result, status = multiply(data.X, data.Y), "success"
	
	default:
			result, status = 0, "failed"
	}

	if status == "success" {

		dataOut := DataOut{result, "geofferyj", data.OperationType}

		json.NewEncoder(w).Encode(dataOut)
	} else {
		fmt.Fprintf(w, `{"status": "%s", "reason": "%s"}`, "failed", "operation type not supported")
	}

}


func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", getHandler).Methods("GET")
	router.HandleFunc("/", postHandler).Methods("POST")
	http.ListenAndServe(":80", router)
}

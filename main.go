package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting server")
	http.HandleFunc("/GET/", GetHandler)
	http.HandleFunc("/POST", PostHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Server")
}

type name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	n := name{}

	fmt.Println(n)

	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(n)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, "Welcome to the Server")
}

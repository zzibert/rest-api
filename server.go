package main

import (
	"database/sql"
	"net/http"
)

func main() {

	var err error
	db, err := sql.Open("postgres", "user=zzibert dbname=postgres password=nekineki port=5432 sslmode=disable")

	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/group/", handleGroupRequest)
	http.HandleFunc("/user/", handleUserRequest)
	server.ListenAndServe()
}

// GROUP HANDLER FUNCTIONS

func handleGroupRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGroupGet(w, r)
	case "POST":
		err = handleGroupPost(w, r)
	case "PUT":
		err = handleGroupPut(w, r)
	case "DELETE":
		err = handleGroupDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// USER HANDLER FUNCTIONS

func handleUserRequest(w http.ResponseWriter, r *http.Request) {

}

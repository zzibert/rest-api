package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"
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

func handleGroupGet(w http.ResponseWriter, r *http.Request) (err error) {
	group := Group{}
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	_, err = group.fetch(id)
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&group, "", "\t\t")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// USER HANDLER FUNCTIONS

func handleUserRequest(w http.ResponseWriter, r *http.Request) {

}

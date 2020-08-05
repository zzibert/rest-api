package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	. "github.com/zzibert/rest-api/data"
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
	http.HandleFunc("/group/", handleGroupRequest(&Group{Db: db}))
	http.HandleFunc("/user/", handleUserRequest(&User{Db: db}))
	server.ListenAndServe()
}

// GROUP HANDLER FUNCTIONS

func handleGroupRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    var err error
    switch r.Method {
    case "GET":
      err = handleGroupGet(w, r, t)
    case "POST":
      err = handleGroupPost(w, r, t)
    case "PUT":
      err = handleGroupPut(w, r, t)
    case "DELETE":
      err = handleGroupDelete(w, r, t)
    }
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
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

func handleGroupPost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var group Group

	json.Unmarshal(body, &group)
	err = group.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleGroupPut(w http.ResponseWriter, r *http.Request) (err error) {
  id, err := strconv.Atoi(path.Base(r.URL.Path))
  if err != nil {
    return
  }

  group, err
}

// USER HANDLER FUNCTIONS

func handleUserRequest(w http.ResponseWriter, r *http.Request) {

}

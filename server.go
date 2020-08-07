package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	_ "github.com/lib/pq"
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
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// USER HANDLER FUNCTIONS

func handleUserRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

//GENERIC HANDLER FUNCTIONS

func handleGet(w http.ResponseWriter, r *http.Request, text Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	err = text.Fetch(id)
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&text, "", "\t\t")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handleGetAll(w http.ResponseWriter, r *http.Request, text Text) (err error) {

	texts, err := text.List()
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&texts, "", "\t\t")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request, text Text) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	json.Unmarshal(body, &text)
	err = text.Create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request, text Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	err = text.Fetch(id)
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &text)
	err = text.Update()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request, text Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	err = text.Fetch(id)
	if err != nil {
		return
	}

	err = text.Delete()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

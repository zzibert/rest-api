// Package classification of API

package api

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	. "github.com/zzibert/rest-api/data"
)

// GROUP HANDLER FUNCTIONS

func HandleGroupRequest(g GroupType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			switch path.Base(r.URL.Path) {
			case "group":
				err = handleGetAllGroups(w, r, g)
			default:
				err = handleGet(w, r, g)
			}
		case "POST":
			err = handlePost(w, r, g)
		case "PUT":
			err = handlePut(w, r, g)
		case "DELETE":
			err = handleDelete(w, r, g)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleGetAllGroups(w http.ResponseWriter, r *http.Request, group GroupType) (err error) {

	groups, err := group.List()
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&groups, "", "\t\t")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// USER HANDLER FUNCTIONS

func HandleUserRequest(u UserType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			switch path.Base(r.URL.Path) {
			case "user":
				err = handleGetAllUsers(w, r, u)
			default:
				err = handleGet(w, r, u)
			}
		case "POST":
			err = handlePost(w, r, u)
		case "PUT":
			err = handlePut(w, r, u)
		case "DELETE":
			err = handleDelete(w, r, u)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleGetAllUsers(w http.ResponseWriter, r *http.Request, user UserType) (err error) {

	users, err := user.List()
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&users, "", "\t\t")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
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

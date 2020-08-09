package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/zzibert/rest-api/data"
)

func TestHandleGroupGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/group/", handleGroupRequest(&TestGroup{}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/group/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var group Group
	json.Unmarshal(writer.Body.Bytes(), &group)
	if group.Id != 1 {
		t.Errorf("Cannot retrieve JSON group")
	}
}

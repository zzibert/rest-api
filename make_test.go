package main

import (
	"net/http"
	"testing"
)

func TestHandleGroupGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/group/", handleGroupRequest)

}

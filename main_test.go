package main

import (
	"net/http"
	"testing"

	. "github.com/zzibert/rest-api/data"
)

func TestHandleGroupGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/group/", handleGroupRequest(&TestGroup{}))
}

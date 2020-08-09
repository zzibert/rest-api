package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/zzibert/rest-api/data"
	. "gopkg.in/check.v1"
)

type GroupTestSuite struct{}

func init() {
	Suite(&GroupTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *GroupTestSuite) TestHandleGet(c *C) {
	mux := http.NewServeMux()
	mux.HandleFunc("/group/", handleGroupRequest(&TestGroup{}))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	c.Check(writer.Code, Equals, 200)
	var group Group
	json.Unmarshal(writer.Body.Bytes(), &group)
	c.Check(group.Id, Equals, 1)
}

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/zzibert/rest-api/data"
	. "gopkg.in/check.v1"
)

type GroupTestSuite struct {
	mux    *http.ServeMux
	group  *TestGroup
	writer *httptest.ResponseRecorder
}

func init() {
	Suite(&GroupTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *GroupTestSuite) SetUpTest(c *C) {
	s.group = &TestGroup{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/group/", handleGroupRequest(s.group))
}

func (s *GroupTestSuite) TestHandleGet(c *C) {
	request, _ := http.NewRequest("GET", "/group/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	var group Group
	json.Unmarshal(s.writer.Body.Bytes(), &group)
	c.Check(group.Id, Equals, 1)
}

func (s *GroupTestSuite) TestHandlePut(c *C) {
	json := strings.NewReader(`{"name":"updated group"}`)
	request, _ := http.NewRequest("PUT", "/group/1", json)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	c.Check(s.group.Id, Equals, 1)
	c.Check(s.group.Name, Equals, "updated group")
}

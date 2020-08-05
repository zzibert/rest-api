package data

import "database/sql"

type User struct {
	Db       *sql.DB
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Group struct {
	Db   *sql.DB
	Id   int    `json:"id"`
	Name string `json:"name"`
}

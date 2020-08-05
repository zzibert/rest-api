package data

import "database/sql"

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

type User struct {
	Db       *sql.DB
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Group    *Group `json:"group"`
}

type Group struct {
	Db    *sql.DB
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}

// GROUP FUNCTIONS

// USER FUNCTIONS

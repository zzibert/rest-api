package data

import (
	"database/sql"
	"errors"
)

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

func (group *Group) fetch(id int) (err error) {
	group.Users = []User{}

	err = group.Db.QueryRow("select id, name from groups where id = $1", id).Scan(&group.Id, &group.Name)
	if err != nil {
		return
	}

	rows, err := group.Db.Query("select id, name, password, email from users where group_id = $1", group.Id)
	if err != nil {
		return
	}

	for rows.Next() {
		user := User{Db: group.Db, Group: group}
		err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		if err != nil {
			return
		}
		group.Users = append(group.Users, user)
	}
	rows.Close()
	return
}

func (group *Group) create() (err error) {
	err = group.Db.QueryRow("insert into groups (name) values ($1) returning id", group.Name).Scan(&group.Id)

	return
}

func (group *Group) update() (err error) {
	_, err = group.Db.Exec("update groups set name = $2 where id = $1", group.Id, group.Name)
	return
}

func (group *Group) delete() (err error) {
	_, err = group.Db.Exec("delete from groups where id = $1", group.Id)
	return
}

// USER FUNCTIONS

func (user *User) create() (err error) {
	if user.Group == nil {
		err = errors.New("Group not found!")
		return
	}
	err = user.Db.QueryRow("insert into users (name, password, email, group_id) values ($1, $2, $3, $4) returning id", user.Name, user.Password, user.Email, user.Group.Id).Scan(&user.Id)
	return
}

func (user *user) fetch(id int) (err error) {

}

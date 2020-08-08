package data

import (
	"database/sql"
	"errors"
)

type Text interface {
	Fetch(id int) (err error)
	Create() (err error)
	Update() (err error)
	Delete() (err error)
}

type User struct {
	Db       *sql.DB
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Group_id int    `json:"group"`
}

type Group struct {
	Db    *sql.DB
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}

// GROUP METHODS

func (group *Group) List() (groups []Group, err error) {
	return ListGroups(group)
}

func ListGroups(group *Group) (groups []Group, err error) {

	rows, err := group.Db.Query("select id, name from groups")
	if err != nil {
		return
	}

	for rows.Next() {
		group := Group{}
		err = rows.Scan(&group.Id, &group.Name)
		if err != nil {
			return
		}
		groups = append(groups, group)
	}
	rows.Close()
	return
}

func (group *Group) Fetch(id int) (err error) {
	group.Users = []User{}

	err = group.Db.QueryRow("select id, name from groups where id = $1", id).Scan(&group.Id, &group.Name)
	if err != nil {
		return
	}

	rows, err := group.Db.Query("select id, name, password, email from users where group_id = $1", id)
	if err != nil {
		return
	}

	for rows.Next() {
		user := User{Db: group.Db, Group_id: id}
		err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		if err != nil {
			return
		}
		group.Users = append(group.Users, user)
	}
	rows.Close()
	return
}

func (group *Group) Create() (err error) {
	err = group.Db.QueryRow("insert into groups (name) values ($1) returning id", group.Name).Scan(&group.Id)

	return
}

func (group *Group) Update() (err error) {
	_, err = group.Db.Exec("update groups set name = $2 where id = $1", group.Id, group.Name)
	return
}

func (group *Group) Delete() (err error) {
	_, err = group.Db.Exec("delete from groups where id = $1", group.Id)
	return
}

// USER METHODS

func (user *User) List() (users []User, err error) {
	return ListUsers(user)
}

func ListUsers(user *User) (users []User, err error) {
	rows, err := user.Db.Query("select id, name, password, email, group_id from users")
	if err != nil {
		return
	}

	users = make([]User, 0)

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Group_id)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

func (user *User) Create() (err error) {

	_, err = user.Db.Exec("select * from groups where id = $1", user.Group_id)
	if err != nil {
		err = errors.New("Group not found!")
		return
	}

	err = user.Db.QueryRow("insert into users (name, password, email, group_id) values ($1, $2, $3, $4) returning id", user.Name, user.Password, user.Email, user.Group_id).Scan(&user.Id)
	return
}

func (user *User) Fetch(id int) (err error) {
	err = user.Db.QueryRow("select id, name, password, email, group_id from users where id = $1", id).Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Group_id)
	return
}

func (user *User) Update() (err error) {
	_, err = user.Db.Exec("update users set name = $2, password = $3, email = $4, group_id = $5 where id = $1", user.Id, user.Name, user.Password, user.Email, user.Group_id)
	return
}

func (user *User) Delete() (err error) {
	_, err = user.Db.Exec("delete from users where id = $1", user.Id)
	return
}

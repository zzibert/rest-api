package data

type TestUser struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Group_id int    `json:"group"`
}

type TestGroup struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}

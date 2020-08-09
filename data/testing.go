package data

// USER

type TestUser struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Group_id int    `json:"group"`
}

func (user *TestUser) Fetch(id int) (err error) {
	user.Id = id
	return
}

func (user *TestUser) Create() (err error) {
	return
}

func (user *TestUser) Update() (err error) {
	return
}

func (user *TestUser) Delete() (err error) {
	return
}

func (user *TestUser) List() (users []UserType, err error) {
	return
}

// GROUP

type TestGroup struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}

func (group *TestGroup) Fetch(id int) (err error) {
	group.Id = id
	return
}

func (group *TestGroup) Create() (err error) {
	return
}

func (group *TestGroup) Update() (err error) {
	return
}

func (group *TestGroup) Delete() (err error) {
	return
}

func (group *TestGroup) List() (groups []GroupType, err error) {
	return
}

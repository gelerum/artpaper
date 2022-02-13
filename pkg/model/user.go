package model

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type GetUser struct {
	Name string `json:"name"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type FindUsers struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

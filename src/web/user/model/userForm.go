package model

type UserForm struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	CheckCode string `json:"checkCode"`
}

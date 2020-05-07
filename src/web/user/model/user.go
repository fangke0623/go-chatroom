package model

type User struct {
	Id         string `json:"id" db:"id" `
	Email      string `json:"email" db:"email"`
	UserName   string `json:"userName" db:"user_name"`
	Password   string `json:"password" db:"password"`
	CreateTime string `json:"createTime" db:"create_time"`
}

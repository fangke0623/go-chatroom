package user

import "wechat/src/common/query"

type Form struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	CheckCode string `json:"checkCode"`
	query.Page
}

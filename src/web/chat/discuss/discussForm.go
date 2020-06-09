package discuss

import "wechat/src/common/query"

type Form struct {
	DiscussId    string `json:"discussId"`
	DiscussTitle string `json:"discussTitle"`
	UserId       string `json:"userId" db:"user_id"`
	VisibleType  string `json:"visibleType"`
	Status       string `json:"status"`
	CreateTime   string `json:"createTime"`
	ModifyTime   string `json:"modifyTime"`
	ModifyId     string `json:"modifyId"`
	query.Page
}

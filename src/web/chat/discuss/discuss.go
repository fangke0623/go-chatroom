package discuss

import "wechat/src/common/query"

type Discuss struct {
	DiscussId    int64  `json:"discussId" db:"discuss_id"`
	DiscussTitle string `json:"discussTitle" db:"discuss_title"`
	UserId       string `json:"userId" db:"user_id"`
	VisibleType  string `json:"visibleType" db:"visible_type"`
	Status       string `json:"status" db:"status"`
	CreateTime   string `json:"createTime" db:"create_time"`
	ModifyTime   string `json:"modifyTime" db:"modify_time"`
	ModifyId     string `json:"modifyId" db:"modify_id"`
	query.Page
}

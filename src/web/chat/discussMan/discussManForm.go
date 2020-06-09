package discussMan

import "wechat/src/common/query"

type Form struct {
	ManId     string `json:"manId" db:"man_id"`
	DiscussId string `json:"discussId" db:"discuss_id"`
	UserId    string `json:"userId" db:"user_id"`
	ManType   string `json:"manType" db:"man_type"`
	Remind    string `json:"remind" db:"remind"`
	Status    string `json:"status" db:"status"`
	query.Page
}

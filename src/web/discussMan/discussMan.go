package discussMan

type DiscussMan struct {
	ManId      int    `json:"manId" db:"man_id"`
	DiscussId  string `json:"discussId" db:"discuss_id"`
	UserId     string `json:"userId" db:"user_id"`
	ManType    string `json:"manType" db:"man_type"`
	Remind     string `json:"remind" db:"remind"`
	Status     string `json:"status" db:"status"`
	CreateTime string `json:"createTime" db:"create_time"`
	ModifyTime string `json:"modifyTime" db:"modify_time"`
}

package discussMan

type Form struct {
	ManId     int    `json:"manId" db:"man_id"`
	DiscussId int    `json:"discussId" db:"discuss_id"`
	MemberId  string `json:"memberId" db:"member_id"`
	ManType   byte   `json:"manType" db:"man_type"`
	Remind    byte   `json:"remind" db:"remind"`
	Status    byte   `json:"status" db:"status"`
}

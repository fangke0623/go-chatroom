package discussMsg

type Form struct {
	MsgId      int    `json:"msgId" db:"msg_id"`
	ManId      int    `json:"manId" db:"man_id"`
	DiscussId  int    `json:"discussId" db:"discuss_id"`
	MsgContent string `json:"msgContent" db:"msg_content"`
	MsgType    byte   `json:"msgType" db:"msg_type"`
	Remind     byte   `json:"remind" db:"remind"`
	Status     byte   `json:"status" db:"status"`
	CreateDate string `json:"createDate" db:"create_date"`
	UpdateDate string `json:"updateDate" db:"update_date"`
}

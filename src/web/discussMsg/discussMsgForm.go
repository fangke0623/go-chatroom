package discussMsg

type Form struct {
	MsgId      string `json:"msgId" db:"msg_id"`
	ManId      string `json:"manId" db:"man_id"`
	DiscussId  string `json:"discussId" db:"discuss_id"`
	MsgContent string `json:"msgContent" db:"msg_content"`
	MsgType    string `json:"msgType" db:"msg_type"`
	Remind     string `json:"remind" db:"remind"`
	Status     string `json:"status" db:"status"`
	CreateDate string `json:"createDate" db:"create_date"`
	UpdateDate string `json:"updateDate" db:"update_date"`
}

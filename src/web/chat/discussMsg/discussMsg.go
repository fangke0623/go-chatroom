package discussMsg

type DiscussMsg struct {
	MsgId      int    `json:"msgId" db:"msg_id"`
	ManId      int    `json:"manId" db:"man_id"`
	DiscussId  int64  `json:"discussId" db:"discuss_id"`
	MsgContent string `json:"msgContent" db:"msg_content"`
	MsgType    int    `json:"msgType" db:"msg_type"`
	Status     string `json:"status" db:"status"`
	CreateTime string `json:"createTime" db:"create_time"`
	ModifyTime string `json:"modifyTime" db:"modify_time"`
}

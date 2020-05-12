package discuss

type Discuss struct {
	DiscussId    int    `json:"discussId" db:"discuss_id"`
	DiscussTitle string `json:"discussTitle" db:"discuss_title"`
	UserId       string `json:"userId" db:"user_id"`
	VisibleType  byte   `json:"visibleType" db:"visible_type"`
	Status       byte   `json:"status" db:"status"`
	CreateDate   string `json:"createDate" db:"create_date"`
	UpdateDate   string `json:"updateDate" db:"update_date"`
	UpdateId     string `json:"updateId" db:"update_id"`
}

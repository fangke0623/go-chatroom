package user

type Form struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	NickName  string `json:"nickName"`
	CheckCode string `json:"checkCode"`
}

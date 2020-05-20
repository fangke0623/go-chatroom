package cookie

import (
	"net/http"
	"wechat/src/web/user"
)

func SetCookie(w http.ResponseWriter, param interface{}) {
	u, _ := param.(user.User)
	cookie := http.Cookie{Name: u.UserName, Value: u.Nickname, Expires: u.LoginTime}

	http.SetCookie(w, &cookie)
}

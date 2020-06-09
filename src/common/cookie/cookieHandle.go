package cookie

import (
	"encoding/json"
	"net/http"
	"wechat/src/config"
	"wechat/src/web/chat/user"
)

func SetCookie(w http.ResponseWriter, param interface{}) {
	u, _ := param.(user.User)

	cookie := http.Cookie{Name: u.Id, Value: u.UserName}
	value, _ := json.Marshal(u)
	config.DoSet(u.Id, value)
	config.DoExpire(u.Id, 0.5*3600)

	http.SetCookie(w, &cookie)
}

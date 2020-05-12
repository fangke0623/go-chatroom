package handle

import (
	"encoding/json"
	"log"
	"net/http"
	"wechat/src/web/discuss"
	"wechat/src/web/discussMan"
	"wechat/src/web/discussMsg"
	"wechat/src/web/user"
)

func ResponseHandle(writer http.ResponseWriter, request *http.Request) {

	param, err := handleParams(request)
	if err != nil {
		log.Print("form转换json异常：", err)
	}
	switch request.URL.Path {
	case "/user/list":
		writer.Write(user.FindUserList(param))
	case "/user/register":
		writer.Write(user.RegisterUser(param))
	case "/user/login":
		writer.Write(user.Login(param))

	case "/discuss/add":
		writer.Write(discuss.AddDiscuss(param))
	case "/discuss/update":
		writer.Write(discuss.UpdateDiscuss(param))
	case "/discuss/delete":
		writer.Write(discuss.DeleteDiscuss(param))
	case "/discuss/list":

		writer.Write(discuss.FindDiscussList(param))
	case "/discussMan/list":
		writer.Write(discussMan.FindDiscussManList(param))
	case "/discussMan/add":
		writer.Write(discussMan.AddDiscussMan(param))
	case "/discussMan/update":
		writer.Write(discussMan.AddDiscussMan(param))

	case "/discussMsg/list":
		writer.Write(discussMsg.FindDiscussMsgList(param))
	case "/discussMsg/add":
		writer.Write(discussMsg.AddDiscussMsg(param))
	}

}
func handleParams(request *http.Request) ([]byte, error) {
	var param []byte
	var err error
	_ = request.ParseForm()
	if request.Method == "POST" {
		param, err = json.Marshal(request.PostForm)
	}
	if request.Method == "GET" {
		param, err = json.Marshal(request.Form)
	}
	return param, err
}

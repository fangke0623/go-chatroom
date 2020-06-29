package handle

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"wechat/src/common/enum"
	"wechat/src/common/exception"
	"wechat/src/common/fileHandle"
	"wechat/src/web/chat/discuss"
	"wechat/src/web/chat/discussMan"
	"wechat/src/web/chat/discussMsg"
	"wechat/src/web/chat/user"
)

func ResponseHandle(writer http.ResponseWriter, request *http.Request) {

	param, err := handleParams(request)
	if err != nil {
		log.Print("form转换json异常：", err)
	}
	var result interface{}

	var e exception.Error

	var discussForm discuss.Form

	//user
	if strings.HasPrefix(request.URL.Path, enum.UserService) {
		result, e = user.Handle(request.URL.Path, param)
	} else {
		switch request.URL.Path {

		//discuss
		case "/discuss/add":
			result, e = discussForm.AddDiscuss(param)
			break
		case "/discuss/edit":
			result, e = discussForm.EditDiscuss(param)
			break
		case "/discuss/delete":
			result, e = discussForm.DeleteDiscuss(param)
			break
		case "/discuss/list":
			result, e = discussForm.FindDiscussList(param)
			break
		case "/discuss/detail":
			result, e = discussForm.GetDiscussDetail(param)
			break
		//discussMan
		case "/discussMan/list":
			result, e = discussMan.FindDiscussManList(param)
			break
		case "/discussMan/add":
			result, e = discussMan.AddDiscussMan(param)
			break
		case "/discussMan/edit":
			result, e = discussMan.EditDiscussMan(param)
			break
		case "/discussMan/delete":
			result, e = discussMan.DeleteDiscussMan(param)
			break

		//discussMsg
		case "/discussMsg/list":
			result, e = discussMsg.FindDiscussMsgList(param)
			break
		case "/discussMsg/add":
			result, e = discussMsg.AddDiscussMsg(param)
			break
		//fileHandle
		case "/uploadBase64":
			fileHandle.UploadBase64(param)
			break

		}

	}
	writerJson(writer, result, e)
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

func writerJson(writer http.ResponseWriter, param interface{}, exception exception.Error) {
	resultMap := make(map[string]interface{})
	resultMap["data"] = param
	resultMap["errorCode"] = exception.ErrorCode
	resultMap["errorMsg"] = exception.ErrorMsg
	result, _ := json.Marshal(resultMap)
	_, err := writer.Write(result)
	if err != nil {
		log.Println(err)
	}
}

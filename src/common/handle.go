package common

import (
	"net/http"
	userService "wechat/src/web/user/service"
)

func HandleResponse(writer http.ResponseWriter, request *http.Request) {

	switch request.URL.Path {
	case "/user/list":
		userService.SelectUserList(writer, request)

	}

}

package common

import (
	userService "chatroom/src/web/user/service"
	"net/http"
)

func HandleResponse(writer http.ResponseWriter, request *http.Request) {

	switch request.URL.Path {
	case "/user/list":
		userService.SelectUserList(writer, request)

	}

}

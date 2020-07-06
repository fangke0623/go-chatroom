package base

import "wechat/src/common/exception"

type Service interface {
	Add(param []byte) (interface{}, exception.Error)

	Edit(param []byte) (interface{}, exception.Error)

	Delete(param []byte) (interface{}, exception.Error)

	GetDetail(param []byte) (interface{}, exception.Error)

	FindList(param []byte) (interface{}, exception.Error)
}

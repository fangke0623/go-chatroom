package util

import "reflect"

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)

	return vi.IsNil()

}

package util

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

func MapToStruct(m map[string]interface{}, output interface{}) {

	mapstructure.Decode(m, output)
}
func HandleParamsToStruct(param []byte, object interface{}) {

	var src map[string][]string
	_ = json.Unmarshal(param, &src)

	var dest map[string]interface{}

	m := src
	n := make(map[string]interface{})
	for k := range m {
		n[k] = m[k][0]
	}
	dest = n
	MapToStruct(dest, object)
}

func MapToJson() {

}

func JsonToMap() {

}
func JsonToStruct() {

}

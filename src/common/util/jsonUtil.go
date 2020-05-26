package util

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"log"
)

func MapToStruct(m map[string]interface{}, output interface{}) {

	err := mapstructure.Decode(m, output)
	if err != nil {
		log.Panicln(err)
	}
}
func HandleParamsToStruct(param []byte, object interface{}) {

	var src map[string][]interface{}
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

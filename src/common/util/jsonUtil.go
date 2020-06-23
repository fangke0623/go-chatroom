package util

import (
	"encoding/json"
	"log"
)

func MapToStruct(m map[string]interface{}, output interface{}) {

	result, err := json.MarshalIndent(m, "", "    ")
	err = json.Unmarshal(result, &output)
	if err != nil {
		log.Panicln(err.Error())
	}
}
func HandleParamsToStruct(param []byte, object interface{}) {

	var src map[string][]interface{}
	err := json.Unmarshal(param, &src)
	if err != nil {
		log.Println(err.Error())
	}

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

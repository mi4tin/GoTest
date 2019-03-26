package jsonHelper

import (
	"encoding/json"
	"moqikaka.com/goutil/logUtil"
)

func JsonString(inputObj interface{}) string {
	bytes, err := json.Marshal(inputObj)
	if err != nil {
		logUtil.ErrorLog("JsonString#出错Log", err)
		return "JsonString#出错"
	}
	return string(bytes)
}
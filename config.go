package goutil

import (
	"encoding/json"
	"io/ioutil"
)

//从文件解析json对象，并映射到强类型对象中,主要用于配置文件解析
func ParseJsonFile(dest interface{}, file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, dest)
	return err
}

//从字符串解析json对象，并映射到强类型对象中
func ParseJsonString(dest interface{}, str string) error {
	err := json.Unmarshal([]byte(str), dest)
	return err
}

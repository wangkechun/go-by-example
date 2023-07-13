package tools

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// BadStrToUtf8 正则替换字符串里的Unicode
func BadStrToUtf8(input string) string {
	reg, _ := regexp.Compile("\\\\u\\w{4}")
	return reg.ReplaceAllStringFunc(input, func(input string) string {
		replaceU := strings.Replace(input, "\\u", "", -1)
		tmp, _ := strconv.ParseInt(replaceU, 16, 32)
		return fmt.Sprintf("%s", string(rune(tmp)))
	})
}

//
func ConvertStruct2UrlCode(Obj interface{}) string {
	jsonMap := ConvertStruct2Map(Obj)
	return ConvertMap2UrlCode(jsonMap)
}

// ConvertStruct2Map 将结构体转换为Map
func ConvertStruct2Map(obj interface{}) map[string]string {
	if obj == nil {
		fmt.Println("cannot convert nil")
		return nil
	}
	data, err := json.Marshal(&obj)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var jsonMap = make(map[string]string)
	err = json.Unmarshal(data, &jsonMap)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return jsonMap
}

// ConvertMap2UrlCode 将map转换为urlCode
func ConvertMap2UrlCode(obj map[string]string) string {
	if obj == nil {
		fmt.Println("cannot convert nil")
		return ""
	}
	var urlCodeStr string
	for k, v := range obj {
		urlCodeStr += fmt.Sprintf("%s=%s&", k, v)
	}
	return urlCodeStr
}

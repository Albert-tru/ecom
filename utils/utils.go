package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

// 解析从前端发送过来的http请求的json数据
func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	//解码到go结构体
	return json.NewDecoder(r.Body).Decode(payload)

}

func WriteJson(w http.ResponseWriter, staus int, v any) error {
	//设置响应头部
	w.Header().Add("Content-Type", "application/json")
	//设置状态码
	w.WriteHeader(staus)
	//编码为json并写入响应体
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, staus int, message string) {
	WriteJson(w, staus, map[string]string{"error": message})

}

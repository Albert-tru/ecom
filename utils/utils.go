package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

//validator.New() 创建的 *validator.Validate 会根据结构体字段的 validate:"..."
//  标签去检查字段值是否满足规则；Validate.Struct(&payload) 调用触发这些检查。

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

// 从请求中提取 token
// 优先从 Authorization 头部获取，其次从 URL 查询参数获取
func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		tokenAuth = strings.TrimSpace(tokenAuth)
		if strings.HasPrefix(tokenAuth, "Bearer ") {
			return strings.TrimSpace(strings.TrimPrefix(tokenAuth, "Bearer "))
		}
		return tokenAuth
	}

	if tokenQuery != "" {
		return strings.TrimSpace(tokenQuery)
	}

	return ""
}

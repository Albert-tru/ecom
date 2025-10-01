package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Albert-tru/ecom/types"
	"github.com/gorilla/mux"
)

/*
测试思路
1. 功能正确：能否成功注册
2. 数据验证：无效数据是否被拒绝
3. 重复注册：相同用户名是否被拒绝
4. 边界情况：极端输入是否被正确处理

*/

// 测试用户服务处理函数
func TestUserServiceHandle(t *testing.T) {

	userStore := &mockUserStore{}    //可控的假仓库
	handler := NewHandler(userStore) //把它注入到“待测的处理器”中

	// 测试用例：成功注册
	t.Run("用户数据无效，注册失败", func(t *testing.T) {
		// 模拟请求数据 - 使用完全无效的数据
		payload := types.RegisterUserPayload{
			Firstname: "JJ",                   // 空firstname
			Lastname:  "DD",                   // 空lastname
			Email:     "invalid-email@qq.com", // 无效邮箱格式
			Password:  "1236",                 // 太短的密码
		}

		// 将请求数据编码为 JSON
		marshalled, _ := json.Marshal(payload)

		// 创建 HTTP 请求							 请求路径
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(marshalled))

		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder() // 创建响应接收器
		router := mux.NewRouter()    // 创建mini路由器

		router.HandleFunc("/register", handler.handleRegister) // 将路径关联到处理函数
		router.ServeHTTP(rr, req)                              // 处理请求

		// 检查响应状态码是否符合预期
		if rr.Code != http.StatusBadRequest {
			t.Errorf("期望状态码 %d, 实际状态码 %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("用户数据有效", func(t *testing.T) {
		// 模拟请求数据
		payload := types.RegisterUserPayload{
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "123@gmail.com", // 故意留空以测试无效数据
			Password:  "123456",
		}

		// 将请求数据编码为 JSON
		marshalled, _ := json.Marshal(payload)

		// 创建 HTTP 请求							 请求路径
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(marshalled))

		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder() // 创建响应接收器
		router := mux.NewRouter()    // 创建mini路由器

		router.HandleFunc("/register", handler.handleRegister) // 将路径关联到处理函数
		router.ServeHTTP(rr, req)                              // 处理请求

		// 检查响应状态码是否符合预期
		if rr.Code != http.StatusCreated {
			t.Errorf("期望状态码 %d, 实际状态码 %d", http.StatusCreated, rr.Code)
		}
	})
}

// 创建一个模拟对象，模仿真实对象的行为但不实际依赖外部系统
type mockUserStore struct {
}

func (m *mockUserStore) CreateUser(user *types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found") // 返回错误表示用户不存在
}
func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

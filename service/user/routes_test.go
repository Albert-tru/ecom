package user

import (
	"bytes"
	"encoding/json"
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
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	// 测试用例：成功注册
	t.Run("用户数据无效，注册失败", func(t *testing.T) {
		// 模拟请求数据
		payload := types.RegisterUserPayload{
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "", // 故意留空以测试无效数据
			Password:  "password123",
		}

		// 将请求数据编码为 JSON
		marshalled, _ := json.Marshal(payload)

		// 创建 HTTP 请求							 请求路径
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		// 创建响应记录器
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		// 注册处理函数
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("期望状态码 %d, 实际状态码 %d", http.StatusBadRequest, rr.Code)
		}
	})
}

// 创建一个模拟对象，模仿真实对象的行为但不实际依赖外部系统
type mockUserStore struct {
}

func (m *mockUserStore) CreateUser(user *types.User) error {
	return nil
}

func (m *mockUserStore) GetUsersByEmail(email string) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) GetUsersByID(id int) (*types.User, error) {
	return nil, nil
}

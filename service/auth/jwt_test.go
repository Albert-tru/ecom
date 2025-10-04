package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Albert-tru/ecom/config"

	"github.com/Albert-tru/ecom/types"
	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateJWT(t *testing.T) {
	// 准备测试数据
	testSecret := []byte(config.Envs.JWTSecret)
	testUserID := 123

	// 调用被测试函数
	tokenString, err := GenerateJWT(testSecret, testUserID)

	// 检查生成是否成功
	if err != nil {
		t.Errorf("GenerateJWT 返回错误: %v", err)
	}

	if tokenString == "" {
		t.Error("GenerateJWT 返回了空字符串")
	}

	// 解析并验证生成的 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return testSecret, nil
	})

	if err != nil {
		t.Errorf("无法解析生成的 token: %v", err)
	}

	if !token.Valid {
		t.Error("生成的 token 无效")
	}

	// 检查 token 中的 claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Error("无法获取 token 的 claims")
	}

	// 验证 user_id claim
	userID, ok := claims["user_id"].(string)
	if !ok {
		t.Error("claims 中没有 user_id 或类型错误")
	}

	if userID != "123" {
		t.Errorf("user_id 期望为 '123', 实际为 '%s'", userID)
	}

	// 验证 exp claim
	exp, ok := claims["exp"].(float64)
	if !ok {
		t.Error("claims 中没有 exp 或类型错误")
	}

	// 确保过期时间在未来
	expTime := time.Unix(int64(exp), 0)
	if expTime.Before(time.Now()) {
		t.Error("生成的 token 已过期")
	}
}

func TestValidateJWT(t *testing.T) {
	// 准备测试数据
	testSecret := []byte(config.Envs.JWTSecret)
	testUserID := 123

	// 生成一个有效的 token
	tokenString, _ := GenerateJWT(testSecret, testUserID)

	// 测试有效的 token
	t.Run("有效的 token", func(t *testing.T) {
		token, err := validateJWT(tokenString)

		if err != nil {
			t.Errorf("validateJWT 返回错误: %v", err)
		}

		if token == nil {
			t.Error("validateJWT 返回了 nil token")
		}

		if !token.Valid {
			t.Error("validateJWT 返回的 token 无效")
		}
	})

	// 测试无效的 token
	t.Run("无效的 token", func(t *testing.T) {
		invalidToken := tokenString + "invalid"

		token, err := validateJWT(invalidToken)

		if err == nil {
			t.Error("validateJWT 没有返回错误")
		}

		if token != nil && token.Valid {
			t.Error("validateJWT 返回了有效的 token")
		}
	})

	// 测试过期的 token
	t.Run("过期的 token", func(t *testing.T) {
		// 创建一个自定义的、已过期的 token
		expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": "123",
			"exp":     time.Now().Add(-24 * time.Hour).Unix(), // 过期时间设为过去
		})

		expiredTokenString, _ := expiredToken.SignedString(testSecret)

		token, err := validateJWT(expiredTokenString)

		if err == nil {
			t.Error("validateJWT 没有为过期的 token 返回错误")
		}

		if token != nil && token.Valid {
			t.Error("validateJWT 认为过期的 token 是有效的")
		}
	})
}

func TestGetTokenFromContext(t *testing.T) {
	// 测试正确的 context
	t.Run("有效的上下文", func(t *testing.T) {
		// 创建带有用户 ID 的上下文
		ctx := context.WithValue(context.Background(), UserKey, 123)

		// 从上下文获取用户 ID
		userID := GetUserIDFromContext(ctx)

		if userID != 123 {
			t.Errorf("GetUserIDFromContext 返回了错误的用户 ID: 期望 123, 实际 %d", userID)
		}
	})

	// 测试无效的 context
	t.Run("无效的上下文", func(t *testing.T) {
		// 创建不包含用户 ID 的上下文
		ctx := context.Background()

		// 在测试中我们期望日志输出，但不想让测试失败
		// 这里可以根据实际代码的行为修改断言
		userID := GetUserIDFromContext(ctx)
		if userID != 0 {
			t.Errorf("GetUserIDFromContext 应该返回 0, 实际返回 %d", userID)
		}
	})
}

func TestGetTokenFromRequest(t *testing.T) {
	// 使用 utils.GetTokenFromRequest 测试，如果有必要，可以在此处添加测试
	// 但它在不同的包中，可能需要一个单独的测试
}

// 测试中间件
func TestWithJWTAuth(t *testing.T) {
	// 由于此函数依赖于外部 store 和 http 处理，我们可能需要使用模拟对象
	// 这是一个更复杂的测试，需要设置模拟的 http 请求和响应

	t.Run("有效的 JWT", func(t *testing.T) {
		// 创建模拟处理函数
		handlerCalled := false
		mockHandler := func(w http.ResponseWriter, r *http.Request) {
			handlerCalled = true
		}

		// 创建模拟 store
		mockStore := &MockUserStore{
			getUserByIDFunc: func(id int) (*types.User, error) {
				return &types.User{ID: id}, nil
			},
		}

		// 创建带有 JWT 的请求
		req, _ := http.NewRequest("GET", "/test", nil)
		token, _ := GenerateJWT([]byte(config.Envs.JWTSecret), 123)
		req.Header.Set("Authorization", "Bearer " + token) // 假设 validateJWT 已被模拟

		// 创建响应记录器
		rr := httptest.NewRecorder()

		// 调用中间件
		handler := WithJWTAuth(http.HandlerFunc(mockHandler), mockStore)
		handler(rr, req)

		// 验证原始处理函数是否被调用
		if !handlerCalled {
			t.Error("原始处理函数没有被调用")
		}
	})

	// 可以添加更多测试场景，如无效的 JWT、无法从 store 获取用户等
}

// MockUserStore 是一个模拟的 UserStore 实现，用于测试
type MockUserStore struct {
	getUserByIDFunc func(id int) (*types.User, error)
	// 可以根据需要添加更多方法
}

func (m *MockUserStore) GetUserByID(id int) (*types.User, error) {
	if m.getUserByIDFunc != nil {
		return m.getUserByIDFunc(id)
	}
	return nil, nil
}

// 实现接口所需的其他方法
func (m *MockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *MockUserStore) CreateUser(u *types.User) error {
	return nil
}

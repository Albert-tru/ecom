package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Albert-tru/ecom/config"
	"github.com/Albert-tru/ecom/types"
	"github.com/Albert-tru/ecom/utils"
	"github.com/golang-jwt/jwt/v5"
)

// UserKey is the key used to store the user ID in context
type contextKey string

const UserKey contextKey = "user_id"

// 用来签名的密钥	 要写进token的用户id
func GenerateJWT(secret []byte, userID int) (string, error) {
	//设置过期时间
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationSeconds)

	// 创建 token				生成签名				map形式存储载荷
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": strconv.Itoa(userID),
		"exp":     time.Now().Add(expiration).Unix(), // 过期时间戳
	})

	// 生成并返回签名字符串
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 验证 JWT 并返回解析后的 token 对象
func validateJWT(tokenString string) (*jwt.Token, error) {
	// 解析 token
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保使用的是预期的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回用于验证签名的密钥
		return []byte(config.Envs.JWTSecret), nil
	})
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 从请求中获取token
		tokenString := utils.GetTokenFromRequest(r)

		// 2. 验证token
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		// 3. 提取用户ID
		claims := token.Claims.(jwt.MapClaims)
		str := claims["user_id"].(string)
		userID, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert userID to int: %v", err)
			permissionDenied(w)
			return
		}

		// 4. 验证用户是否存在
		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		// 5. 将用户ID存入Context
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		// 6. 执行实际的处理函数
		handlerFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	http.Error(w, "Permission denied", http.StatusForbidden)
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return 0
	}

	return userID
}

func GetTokenFromRequest(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if token == "" {
		// 尝试从查询参数中获取 token
		token = r.URL.Query().Get("token")
	}

	return token
}

func validateToken(tokenString string) (*jwt.Token, error) {
	// 解析 token
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保使用的是预期的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回用于验证签名的密钥
		return []byte(config.Envs.JWTSecret), nil
	})
}

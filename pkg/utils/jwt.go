package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	// ExpireDuration 20秒
	ExpireDuration = 20 * time.Second
	JwtSecretKey   = `ihawe98fu98234auth2`
)

type MyClaims struct {
	jwt.RegisteredClaims
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"` // 自定义声明（例如角色）
}

// GenerateToken 生成token
func GenerateToken(id int64, username string) (string, error) {
	// 定义 token 的过期时间
	expireTime := time.Now().Add(ExpireDuration).Unix()
	// 使用 JWT 签名算法生成 token
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "my-issuer",                                  // 签发者
			Subject:   username,                                     // 主题（用户ID）
			ExpiresAt: jwt.NewNumericDate(time.Unix(expireTime, 0)), // 过期时间
		},
		Id:       id,
		Username: username,
		Role:     "admin", // 自定义声明（例如角色）
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 将 token 进行加盐加密
	tokenString, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析 token 并检查是否过期
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 返回密钥进行验证
		return []byte(JwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorClaimsInvalid)
	}
}

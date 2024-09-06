package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	// ExpireDuration 1小时
	ExpireDuration = 3600 * time.Second
	JwtSecretKey   = `ihawe98fu98234auth2`
)

type MyClaims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	jwt.MapClaims
}

// GenerateToken 生成token
func GenerateToken(id int64, username string) (string, error) {
	// 定义token的过期时间
	expireTime := time.Now().Add(ExpireDuration).Unix()
	// 使用 JWT 签名算法生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "my-issuer", // 签发者
		"sub":  username,    // 主题（用户ID）
		"exp":  expireTime,  // 过期时间
		"role": "admin",     // 自定义声明（例如角色）
	})
	// 将token进行加盐加密
	// 注意：该方法参数虽然是interface{}，但是必须传入[]byte类型的参数
	tokenString, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil
}

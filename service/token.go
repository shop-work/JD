/*******
* @Author:qingmeng
* @Description:
* @File:token
* @Date2022/2/16
 */

package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"shop/model"
	"time"
)

var Secret = []byte("人生路漫漫")

// CreateToken 创建token
func CreateToken(user model.User, ExpireTime int64, tokenType string) (string, error) {
	cla := model.MyClaims{
		User: user,
		Type: tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + ExpireTime, //过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)

	return token.SignedString(Secret) // 进行签名生成对应的token
}

// ParseToken 解析token
func ParseToken(tokenString string) (*model.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

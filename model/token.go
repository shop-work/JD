/*******
* @Author:qingmeng
* @Description:
* @File:token
* @Date2022/2/16
 */

package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	User User
	Type string //"REFRESH_TOKEN"表示为一个refresh token，"TOKEN"表示为一个token
	Time time.Time
	jwt.StandardClaims
}

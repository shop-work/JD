/*******
* @Author:qingmeng
* @Description:
* @File:middleware
* @Date2021/12/10
 */

package api

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/service"
	"shop/tool"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

/* cookie验证
func auth(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		tool.RespErrorWithData(ctx, "请先登录")
		ctx.Abort()
	}
	ctx.Set("iUsername", username)
	ctx.Next()
}*/

//jwt验证
func jwtAuth(ctx *gin.Context) {
	token := ctx.PostForm("token")
	if token == "" {
		tool.RespSuccessfulWithData(ctx, "token为空")
		ctx.Abort()
		return
	}

	// 解析到控制台
	jwt := strings.Split(token, ".")
	cnt := 0
	for _, val := range jwt {
		cnt++
		if cnt == 3 {
			break
		}
		msg, _ := base64.StdEncoding.DecodeString(val)
		fmt.Println("val ->", string(msg))
	}
	mc, err := service.ParseToken(token)
	if err != nil {
		fmt.Println("jwtAuthErr:", err.Error())
		tool.RespSuccessfulWithData(ctx, "token无效")
		ctx.Abort()
		return
	}
	ctx.Set("iUsername", mc.User.Username)
	ctx.Next()
}

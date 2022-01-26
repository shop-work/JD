/*******
* @Author:qingmeng
* @Description:
* @File:resp
* @Date2021/12/10
 */

package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespErrorWithData(ctx *gin.Context,data interface{})  {
	ctx.JSON(http.StatusOK,gin.H{
		"info":data,
	})
}

func RespInternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"info": "服务器错误",
	})
}

func RespSuccessful(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
	})
}

func RespSuccessfulWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
		"data": data,
	})
}

func RespSensitiveError(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,gin.H{
		"data":"含有非法词汇",
	})
}
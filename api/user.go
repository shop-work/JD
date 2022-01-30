/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shop/model"
	"shop/service"
	"shop/tool"
	"strconv"
)

//注册
func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//检验用户名是否含有敏感词
	flag := tool.CheckIfSensitive(username)
	if flag {
		tool.RespSensitiveError(ctx)
		return
	}

	user := model.User{
		Username: username,
		Password: password,
	}
	u := service.UserService{}
	flag, err := u.IsExistUsername(username)
	if err != nil {
		fmt.Println("judge exist username err:", err)
		tool.RespInternalError(ctx)
		return
	}

	if flag {
		tool.RespErrorWithData(ctx, "用户名已存在")
		return
	}
	flag = u.IsPasswordReasonable(password)
	if !flag {
		tool.RespErrorWithData(ctx, "密码不合理")
		return
	}

	err = u.Register(user)
	if err != nil {
		fmt.Println("judge repeat username err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "注册成功")
}

//登陆
func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	u := service.UserService{}

	flag, err := u.IsPasswordCorrect(username, password)
	if err != nil {
		fmt.Println("judge password correct err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "密码错误")
		return
	}

	//设置cookie
	ctx.SetCookie("username", username, 3600, "/", "", false, false)

	tool.RespSuccessfulWithData(ctx, "欢迎您~"+username)
}

//给某用户充值
func addMoney(ctx *gin.Context) {
	username := ctx.PostForm("username")
	money := ctx.PostForm("money")
	fMoney, err := strconv.ParseFloat(money, 32)
	if err != nil {
		fmt.Println("money string to float err:", err)
		tool.RespInternalError(ctx)
		return
	}
	u := service.UserService{}
	err = u.AddMoney(username, float32(fMoney))
	if err != nil {
		fmt.Println("addMoney err:", err)
		tool.RespInternalError(ctx)
	}

	tool.RespErrorWithData(ctx, "成功充值"+money+"元")
}

//完善信息
func changeInformation(ctx *gin.Context) {
	iUsername, _ := ctx.Get("iUsername")
	username := iUsername.(string)
	u := service.UserService{}

	//添加电话
	phone := ctx.PostForm("phone")
	if phone != "" {
		err := u.AddPhone(username, phone)
		if err != nil {
			fmt.Println("AddPhone err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	//添加昵称
	//添加收货地址
	tool.RespSuccessful(ctx)
}

//查看个人信息
func viewUserInfo(ctx *gin.Context) {
	iUsername, _ := ctx.Get("iUsername")
	username := iUsername.(string)
	u := service.UserService{}

	userinfo, err := u.GetUserinfo(username)
	if err != nil {
		fmt.Println("getUserinfo err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//借鉴达达子
	outPutMap := tool.ObjToMap(userinfo)
	tool.RespSuccessfulWithData(ctx, outPutMap)

}

//查看余额
func viewUserMoney(ctx *gin.Context) {
	iUsername, _ := ctx.Get("iUsername")
	username := iUsername.(string)
	u := service.UserService{}

	userinfo, err := u.GetUserinfo(username)
	if err != nil {
		fmt.Println("getUserinfo err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, userinfo.Money)

}

//登陆后修改密码
func changePassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")
	iUsername, _ := ctx.Get("iUsername")
	username := iUsername.(string)
	u := service.UserService{}
	//检验旧密码是否正确
	flag, err := u.IsPasswordCorrect(username, oldPassword)
	if err != nil {
		fmt.Println("judge password correct err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "旧密码错误")
		return
	}

	//验证新密码合理性
	flag = u.IsPasswordReasonable(newPassword)
	if !flag {
		tool.RespErrorWithData(ctx, "新密码不合理")
		return
	}

	//修改为新密码
	err = u.ChangePassword(username, newPassword)
	if err != nil {
		fmt.Println("change password err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

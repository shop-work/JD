/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"shop/model"
	"shop/service"
	"shop/tool"
	"strconv"
	"time"
)

var conf = model.Conf{
	ClientId:     "ee9d6a9dad9fced742c1",
	ClientSecret: "abaedca8d689f743fdfd61bd47606abbeecd2c49",
	RedirectUrl:  "http://localhost:8080/api/oauth/redirect",
}

//获取github返回的code
func getCode(ctx *gin.Context) {

	// 解析指定文件生成模板对象
	var temp *template.Template
	var err error
	if temp, err = template.ParseFiles("login.html"); err != nil {
		fmt.Println("读取文件失败，错误信息为:", err)
		tool.RespInternalError(ctx)
		return
	}

	// 利用给定数据渲染模板(html页面)，并将结果写入w，返回给前端
	if err = temp.Execute(ctx.Writer, conf); err != nil {
		fmt.Println("读取渲染html页面失败，错误信息为:", err)
		tool.RespInternalError(ctx)
		return
	}

}

func loginByGithub(ctx *gin.Context) {
	code := ctx.Query("code") //github传来的code

	tokenAuthUrl := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		conf.ClientId, conf.ClientSecret, code)

	fmt.Println(tokenAuthUrl)

	//获取token
	var token *model.Token
	var err error
	if token, err = service.GetToken(tokenAuthUrl); err != nil {
		fmt.Println("get token err:", err)
		tool.RespInternalError(ctx)
		return
	}

	fmt.Printf("%+v", token)

	// 通过token，获取用户信息
	var userInfo map[string]interface{}
	if userInfo, err = service.GetUserInfo(token); err != nil {
		fmt.Println("获取用户信息失败，错误信息为:", err)
		tool.RespInternalError(ctx)
		return
	}

	var userInfoBytes []byte
	if userInfoBytes, err = json.Marshal(userInfo); err != nil {
		fmt.Println("在将用户信息(map)转为用户信息([]byte)时发生错误，错误信息为:", err)
		tool.RespInternalError(ctx)
		return
	}

	var githubUserinfo model.GitHubUserinfo
	err = json.Unmarshal(userInfoBytes, &githubUserinfo)
	if err != nil {
		fmt.Println("unmarshal json err:", err)
		tool.RespInternalError(ctx)
		return
	}

	/*fmt.Println()
	fmt.Println(userInfo)

	ctx.JSON(200,gin.H{
		"userInfo": userInfo,
	})

	fmt.Println(githubUserinfo.Name)*/

	//是否新用户
	u := service.UserService{}
	flag, err := u.IsExistGithubLogin(githubUserinfo.Login)
	if err != nil {
		fmt.Println("judge exist user err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//注册新用户
	if !flag {
		user := model.User{
			Username:    githubUserinfo.Login,
			Name:        githubUserinfo.Name,
			GithubLogin: githubUserinfo.Login,
		}
		err = u.Register(user)
		if err != nil {
			fmt.Println("register err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	//用户登陆
	//获取用户固定信息
	basicUserinfo, err := u.GetBasicUserinfo(githubUserinfo.Login)
	if err != nil {
		fmt.Println("getBasicUserInfo err:", err)
		tool.RespInternalError(ctx)
		return
	}
	//创建token,有效期5分钟
	tokenString, err := service.CreateToken(basicUserinfo, 300, "TOKEN")
	if err != nil {
		fmt.Println("create token err:", err)
		tool.RespInternalError(ctx)
		return
	}
	//创建refreshToken，有效期5天
	refreshTokenString, err := service.CreateToken(basicUserinfo, 5*24*60*60, "TOKEN")
	if err != nil {
		fmt.Println("create token err:", err)
		tool.RespInternalError(ctx)
		return
	}
	fmt.Println(githubUserinfo)

	ctx.JSON(200, gin.H{
		"status":       true,
		"data":         "登陆成功",
		"uid":          basicUserinfo.Uid,
		"groupId":      basicUserinfo.GroupId,
		"token":        tokenString,
		"refreshToken": refreshTokenString,
	})
}

//注册
func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	name := ctx.PostForm("name") //昵称

	//检验用户名是否含有敏感词
	flag := tool.CheckIfSensitive(username)
	flag2 := tool.CheckIfSensitive(name)
	if flag || flag2 {
		tool.RespSensitiveError(ctx)
		return
	}

	u := service.UserService{}

	//用户名是否已存在
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
	if name == "" {
		name = "匿名用户"
	}

	user := model.User{
		Username: username,
		Password: password,
		Salt:     strconv.FormatInt(time.Now().Unix(), 10),
		Name:     name,
	}

	//md5加密
	m5 := md5.New()
	m5.Write([]byte(user.Password))
	m5.Write([]byte(user.Salt))
	st := m5.Sum(nil)
	user.Password = hex.EncodeToString(st)

	//注册
	err = u.Register(user)
	if err != nil {
		fmt.Println("register err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	ctx.JSON(200, gin.H{
		"status":   true,
		"data":     "注册成功",
		"username": username,
		"name":     name,
		"groupId":  0,
	})
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

	//获取用户固定信息
	basicUserinfo, err := u.GetBasicUserinfo(username)
	if err != nil {
		fmt.Println("getBasicUserInfo err:", err)
		tool.RespInternalError(ctx)
		return
	}
	//创建token,有效期5分钟
	tokenString, err := service.CreateToken(basicUserinfo, 300, "TOKEN")
	if err != nil {
		fmt.Println("create token err:", err)
		tool.RespInternalError(ctx)
		return
	}
	//创建refreshToken，有效期5天
	refreshTokenString, err := service.CreateToken(basicUserinfo, 5*24*60*60, "TOKEN")
	if err != nil {
		fmt.Println("create token err:", err)
		tool.RespInternalError(ctx)
		return
	}

	ctx.JSON(200, gin.H{
		"status":       true,
		"data":         "登陆成功",
		"uid":          basicUserinfo.Uid,
		"groupId":      basicUserinfo.GroupId,
		"token":        tokenString,
		"refreshToken": refreshTokenString,
	})

}

//给某用户充值
func addMoney(ctx *gin.Context) {
	//username为空则给自己充值
	username := ctx.PostForm("username")
	if username == "" {
		iUsername, _ := ctx.Get("iUsername")
		username = iUsername.(string)
	}

	money := ctx.PostForm("money")
	if money == "" {
		tool.RespErrorWithData(ctx, "充值金额错误")
		return
	}
	fMoney, err := strconv.ParseFloat(money, 32)
	if err != nil {
		fmt.Println("money string to float err:", err)
		tool.RespErrorWithData(ctx, "充值金额错误")
		return
	}
	if fMoney <= 0 {
		tool.RespErrorWithData(ctx, "充值金额错误")
		return
	}
	u := service.UserService{}
	//是否存在该用户
	flag, err := u.IsExistUsername(username)
	if !flag {
		tool.RespErrorWithData(ctx, "username无效")
		return
	}
	user, err := u.GetUserinfoByUserName(username)
	if err != nil {
		fmt.Println("get userinfo err:", err)
		tool.RespInternalError(ctx)
		return
	}
	err = u.UpdateMoney(username, user.Money+float32(fMoney))
	if err != nil {
		fmt.Println("addMoney err:", err)
		tool.RespInternalError(ctx)
	}

	tool.RespSuccessfulWithData(ctx, "给用户"+username+"成功充值"+money+"元!")
}

//更新信息
func changeInformation(ctx *gin.Context) {
	iUsername, _ := ctx.Get("iUsername")
	username := iUsername.(string)
	u := service.UserService{}

	//更新电话
	phone := ctx.PostForm("phone")

	if phone != "" {
		if len(phone) != 11 {
			tool.RespErrorWithData(ctx, "phone无效")
			return
		}
		err := u.UpdatePhone(username, phone)
		if err != nil {
			fmt.Println("UpdatePhone err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	//更新昵称
	name := ctx.PostForm("name")
	if name != "" {
		flag := tool.CheckIfSensitive(name)
		if flag {
			tool.RespErrorWithData(ctx, "name格式不符合要求")
			return
		}
		err := u.UpdateName(username, name)
		if err != nil {
			fmt.Println("update name err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	//更新性别
	gender := ctx.PostForm("gender")
	if gender != "" {
		err := u.UpdateGender(username, gender)
		if err != nil {
			fmt.Println("update gender err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	//默认选择的收货地址修改
	iAddressId := ctx.PostForm("address_id") //检查该addressId是否有效
	if iAddressId != "" {
		addressId, err := strconv.Atoi(iAddressId)
		if err != nil {
			fmt.Println("addressId to int err:", err)
			tool.RespInternalError(ctx)
			return
		}
		err = u.UpdateAddressId(username, addressId)
		if err != nil {
			fmt.Println("update addressId err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	tool.RespSuccessfulWithData(ctx, "修改成功")
}

//查看个人信息
func viewUserInfo(ctx *gin.Context) {
	iUsername, _ := ctx.Get("iUsername")
	username := iUsername.(string)
	u := service.UserService{}

	userinfo, err := u.GetUserinfoByUserName(username)
	if err != nil {
		fmt.Println("getUserinfo err:", err)
		tool.RespInternalError(ctx)
		return
	}

	outPutMap := tool.ObjToMap(userinfo)
	tool.RespSuccessfulWithData(ctx, outPutMap)

}

//查看余额
func viewUserMoney(ctx *gin.Context) {
	iUsername, _ := ctx.Get("iUsername")
	username := iUsername.(string)
	u := service.UserService{}

	userinfo, err := u.GetUserinfoByUserName(username)
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
		tool.RespErrorWithData(ctx, "新密码无效")
		return
	}

	//修改为新密码
	err = u.ChangePassword(username, newPassword)
	if err != nil {
		fmt.Println("change password err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "修改成功,请重新登陆！")
}

//注册商家
func upgradePower(ctx *gin.Context) {
	iUsername, _ := ctx.Get("iUsername")
	token := ctx.PostForm("token")
	username := iUsername.(string)
	u := service.UserService{}

	//验证是否以为商家身份
	mc, _ := service.ParseToken(token)
	if mc.User.GroupId == 1 {
		tool.RespErrorWithData(ctx, "该用户已是商家")
		return
	}
	err := u.UpdateGroupId(username, 1)
	if err != nil {
		fmt.Println("update group id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "注册成功，请重新登陆！")
}

//加入店铺
func addStoreUser(ctx *gin.Context) {
	token := ctx.PostForm("token")
	storeId := ctx.PostForm("store_id")

	//验证商家
	mc, _ := service.ParseToken(token)
	if mc.User.GroupId != 1 {
		tool.RespErrorWithData(ctx, "该用户不是商家")
		return
	}
	if mc.User.StoreId != 0 {
		tool.RespErrorWithData(ctx, "该用户已有店铺")
		return
	}

	//验证storeId
	if storeId == "" {
		tool.RespErrorWithData(ctx, "store_id无效")
		return
	}
	sid, err := strconv.Atoi(storeId)
	if err != nil {
		fmt.Println("store id to int err:", err)
		tool.RespErrorWithData(ctx, "store_id无效")
		return
	}
	ss := service.StoreService{}
	flag, err := ss.IsExistStoreId(sid)
	if err != nil {
		fmt.Println("judge store id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "store_id无效")
		return
	}

	us := service.UserService{}
	err = us.AddStoreUser(mc.User.Username, sid)
	if err != nil {
		fmt.Println("add store user err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "加入"+storeId+"号店铺成功！")
}

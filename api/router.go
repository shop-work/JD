/*******
* @Author:qingmeng
* @Description:
* @File:router
* @Date2021/12/10
 */

package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()
	engine.Use(Cors())
	engine.POST("/user", register) //注册
	engine.GET("/user", login)     //登陆

	userGroup := engine.Group("/user")
	{

		userGroup.POST("/information", auth, changeInformation) //修改信息
		userGroup.GET("/information", auth, viewUserInfo)       //查看信息
		userGroup.POST("/money", addMoney)                      //充值
		userGroup.GET("/money", viewUserMoney)                  //查看个人余额

		//密码组
		userGroup.POST("/password")
		passwordGroup := userGroup.Group("/password")
		{
			passwordGroup.POST("/", auth, changePassword) //登陆后的直接修改密码
		}
	}
	engine.Run(":8080")
}

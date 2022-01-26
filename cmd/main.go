/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date2021/12/10
 */

package main

import (
	"shop/api"
	"shop/dao"
)

func main() {
	dao.InitDB()
	api.InitEngine()
}

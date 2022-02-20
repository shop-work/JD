/*******
* @Author:qingmeng
* @Description:
* @File:store
* @Date2022/2/19
 */

package model

//店铺
type Store struct {
	StoreId    int     `json:"store_id"` //店铺ID
	StoreName  string  `json:"store_name"`
	Notice     string  `json:"notice"`      //店铺公告
	StoreMoney float64 `json:"store_money"` //店铺资金
}

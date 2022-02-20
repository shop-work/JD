/*******
* @Author:qingmeng
* @Description:
* @File:goods
* @Date2022/2/18
 */

package model

import "time"

type Goods struct {
	GoodsId       int       `json:"goods_id"`
	SortId        int       `json:"sort_id"`        //类别id
	StoreId       int       `json:"store_id"`       //店铺id
	GoodsName     string    `json:"goods_name"`     //商品名字
	Picture       string    `json:"picture"`        //商品图片
	Price         float32   `json:"price"`          //价格
	GoodsIntro    string    `json:"goods_intro"`    //商品介绍
	Number        int       `json:"number"`         //商品数量
	Turnover      string    `json:"turnover"`       //成交量
	CommentNumber int       `json:"comment_number"` //评价数
	FeedBack      float32   `json:"feed_back"`      //好评率
	Style         string    `json:"type"`           //款型
	ShelfDate     time.Time `json:"shelf_date"`     //上架日期
}

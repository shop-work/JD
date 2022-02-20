/*******
* @Author:qingmeng
* @Description:
* @File:comment
* @Date2022/2/20
 */

package model

import "time"

type Comment struct {
	CommentId int       `json:"comment_id"`
	GoodsId   int       `json:"goods_id"`
	Uid       int       `json:"uid"`
	Text      string    `json:"text"` //评论内容
	Star      int       `json:"star"` //评星
	Date      time.Time `json:"date"`
}

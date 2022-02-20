/*******
* @Author:qingmeng
* @Description:
* @File:comment
* @Date2022/2/20
 */

package dao

import (
	"shop/model"
)

type CommentDao struct {
}

// AddComment 添加评论
func (d *CommentDao) AddComment(comment model.Comment) error {
	_, err := DB.Exec("insert into shop.comment (goods_id, uid, text, star, comment_date ) values (?,?,?,?,?);", comment.GoodsId, comment.Uid, comment.Text, comment.Star, comment.Date)
	return err
}

// SelectCommentsByGoodsId 根据goodsId返回comments
func (d *CommentDao) SelectCommentsByGoodsId(goodsId int) ([]model.Comment, error) {
	var comments []model.Comment
	rows, err := DB.Query("select * from shop.comment where goods_id=?", goodsId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		comment := model.Comment{}
		err = rows.Scan(&comment.CommentId, &comment.GoodsId, &comment.Uid, &comment.Text, &comment.Star, &comment.Date)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, err
}

// SelectCommentsByUid 根据uid返回comments
func (d *CommentDao) SelectCommentsByUid(uid int) ([]model.Comment, error) {
	var comments []model.Comment
	rows, err := DB.Query("select * from shop.comment where uid=?", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		comment := model.Comment{}
		err = rows.Scan(&comment.CommentId, &comment.GoodsId, &comment.Uid, &comment.Text, &comment.Star, &comment.Date)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, err
}

// SelectCommentsByUidGoodsId 根据uid和goodsId返回comments
func (d *CommentDao) SelectCommentsByUidGoodsId(uid int, goodsId int) ([]model.Comment, error) {
	var comments []model.Comment
	rows, err := DB.Query("select * from shop.comment where uid=? and goods_id=?", uid, goodsId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		comment := model.Comment{}
		err = rows.Scan(&comment.CommentId, &comment.GoodsId, &comment.Uid, &comment.Text, &comment.Star, &comment.Date)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, err
}

/*******
* @Author:qingmeng
* @Description:
* @File:comment
* @Date2022/2/20
 */

package service

import (
	"database/sql"
	"shop/dao"
	"shop/model"
)

type CommentService struct {
}

func (s *CommentService) AddComment(comment model.Comment) error {
	d := dao.CommentDao{}
	return d.AddComment(comment)
}

func (s *CommentService) SelectCommentsByGoodsId(goodsId int) ([]model.Comment, error) {
	d := dao.CommentDao{}
	return d.SelectCommentsByGoodsId(goodsId)
}

func (s *CommentService) SelectCommentsByUid(uid int) ([]model.Comment, error) {
	d := dao.CommentDao{}
	return d.SelectCommentsByUid(uid)
}

func (s *CommentService) SelectCommentsByUidGoodsId(uid int, goodsId int) ([]model.Comment, error) {
	d := dao.CommentDao{}
	return d.SelectCommentsByUidGoodsId(uid, goodsId)
}

func (s *CommentService) IsExistGoodsId(gid int) (bool, error) {
	d := dao.CommentDao{}
	_, err := d.SelectCommentsByGoodsId(gid)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

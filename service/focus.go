/*******
* @Author:qingmeng
* @Description:
* @File:focus
* @Date2022/2/20
 */

package service

import (
	"database/sql"
	"shop/dao"
	"shop/model"
)

type FocusService struct {
}

func (s *FocusService) IsExistGoodsFocus(uid int, goodsId int) (bool, error) {
	d := dao.FocusDao{}
	_, err := d.SelectFocusByUidGoodsId(uid, goodsId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *FocusService) IsExistStoreFocus(uid int, storeId int) (bool, error) {
	d := dao.FocusDao{}
	_, err := d.SelectFocusByUidStoreId(uid, storeId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *FocusService) SelectFocusByUid(uid int) ([]model.Focus, error) {
	d := dao.FocusDao{}
	return d.SelectFocusByUid(uid)
}

func (s *FocusService) DeleteFocusByUid(uid int) error {
	d := dao.FocusDao{}
	return d.DeleteFocusByUid(uid)
}

func (s *FocusService) AddFocus(focus model.Focus) error {
	d := dao.FocusDao{}
	return d.AddFocus(focus)
}

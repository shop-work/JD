/*******
* @Author:qingmeng
* @Description:
* @File:focus
* @Date2022/2/20
 */

package dao

import "shop/model"

type FocusDao struct {
}

// SelectFocusByUid 查看用户关注
func (d *FocusDao) SelectFocusByUid(uid int) ([]model.Focus, error) {
	var focuses []model.Focus
	rows, err := DB.Query("select * from shop.focus where uid=?", uid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var focus model.Focus
		err = rows.Scan(&focus.Uid, &focus.GoodsId, &focus.StoreId)
		if err != nil {
			return nil, err
		}
		focuses = append(focuses, focus)
	}
	return focuses, err
}

func (d *FocusDao) SelectFocusByUidGoodsId(uid int, goodsId int) (model.Focus, error) {
	var focus model.Focus
	row := DB.QueryRow("select * from shop.focus where uid=? and goods_id=?", uid, goodsId)
	if row.Err() != nil {
		return focus, row.Err()
	}
	err := row.Scan(&focus.Uid, &focus.GoodsId, &focus.StoreId)
	return focus, err
}

func (d *FocusDao) SelectFocusByUidStoreId(uid int, storeId int) (model.Focus, error) {
	var focus model.Focus
	row := DB.QueryRow("select * from shop.focus where uid=? and store_id=?", uid, storeId)
	if row.Err() != nil {
		return focus, row.Err()
	}
	err := row.Scan(&focus.Uid, &focus.GoodsId, &focus.StoreId)
	return focus, err
}

// DeleteFocusByUid 删除库中无用关注
func (d *FocusDao) DeleteFocusByUid(uid int) error {
	_, err := DB.Exec("delete from shop.focus where uid=? and goods_id=0 and store_id=0", uid)
	return err
}

// AddFocus 添加关注
func (d *FocusDao) AddFocus(focus model.Focus) error {
	_, err := DB.Exec("insert into shop.focus (uid, goods_id, store_id ) values (?,?,?);", focus.Uid, focus.GoodsId, focus.StoreId)
	return err
}

/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package dao

import (
	"shop/model"
)

type UserDao struct {
}

func (dao *UserDao) InsertUser(user model.User) error {
	_, err := DB.Exec("insert into shop.userinfo(username, password)values(?,?)", user.Username, user.Password)
	return err
}

//查看用户详细信息
func (dao *UserDao) SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}
	row := DB.QueryRow("select user_id,password,phone,money,address_id from shop.userinfo where username=?", username)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.UserId, &user.Password, &user.Phone, &user.Money, &user.AddressId)
	if err != nil {
		return user, err
	}
	user.Username = username
	return user, err
}

func (dao *UserDao) UpdatePassword(username, newPassword string) error {
	_, err := DB.Exec("update shop.userinfo set password=? where username=?", newPassword, username)
	return err
}

//添加用户电话
func (dao *UserDao) AddPhone(username string, phone string) error {
	_, err := DB.Exec("update shop.userinfo set phone=? where username=?", phone, username)
	return err
}

//充值
func (dao *UserDao) AddMoney(username string, money float32) error {
	_, err := DB.Exec("update shop.userinfo set money=? where username=?", money, username)
	return err
}

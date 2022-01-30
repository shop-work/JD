/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package service

import (
	"database/sql"
	"shop/dao"
	"shop/model"
)

type UserService struct {
}

// Register 注册
func (u *UserService) Register(user model.User) error {
	d := dao.UserDao{}
	err := d.InsertUser(user)
	return err
}

func (u *UserService) IsPasswordCorrect(username, password string) (bool, error) {
	d := dao.UserDao{}
	user, err := d.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	if user.Password != password {
		return false, err
	}
	return true, err
}

// IsExistUsername 判断用户名是否存在
func (u *UserService) IsExistUsername(username string) (bool, error) {
	d := dao.UserDao{}
	_, err := d.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (u *UserService) ChangePassword(username, newPassword string) error {
	d := dao.UserDao{}
	err := d.UpdatePassword(username, newPassword)
	return err
}

// IsPasswordReasonable 验证密码是否合理(可增加密码复杂性)
func (u *UserService) IsPasswordReasonable(password string) bool {
	if len(password) < 6 {
		return false
	}
	return true
}

//添加电话
func (u *UserService) AddPhone(username string, phone string) error {
	d := dao.UserDao{}
	err := d.AddPhone(username, phone)
	return err
}

//充值
func (u *UserService) AddMoney(username string, money float32) error {
	d := dao.UserDao{}
	err := d.AddMoney(username, money)
	return err
}

//获取个人信息
func (u *UserService) GetUserinfo(username string) (user model.User, err error) {
	d := dao.UserDao{}
	user, err = d.SelectUserByUsername(username)
	return user, err
}

//查看余额

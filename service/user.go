/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package service

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/rand"
	"shop/dao"
	"shop/model"
	"time"
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

	//验证密码
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(user.Salt))
	st := m5.Sum(nil)
	password = hex.EncodeToString(st)
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

// UpdatePhone 更新电话
func (u *UserService) UpdatePhone(username string, phone string) error {
	d := dao.UserDao{}
	err := d.UpdatePhone(username, phone)
	return err
}

// GetUserinfoByUserName 根据username获取个人信息
func (u *UserService) GetUserinfoByUserName(username string) (user model.User, err error) {
	d := dao.UserDao{}
	user, err = d.SelectUserByUsername(username)
	return user, err
}

// GetBasicUserinfo 获取用户固定信息
func (u *UserService) GetBasicUserinfo(username string) (user model.User, err error) {
	d := dao.UserDao{}
	user, err = d.SelectBasicUserByUsername(username)
	return user, err
}

// UpdateName 更新昵称
func (u *UserService) UpdateName(username string, name string) error {
	d := dao.UserDao{}
	err := d.UpdateName(username, name)
	return err
}

func (u *UserService) UpdateGender(username string, gender string) error {
	d := dao.UserDao{}
	err := d.UpdateGender(username, gender)
	return err
}

func (u *UserService) UpdateAddressId(username string, id int) error {
	d := dao.UserDao{}
	return d.UpdateAddressId(username, id)
}

func (u *UserService) UpdateGroupId(username string, i int) error {
	d := dao.UserDao{}
	return d.UpdateGroupId(username, i)
}

func (u *UserService) AddStoreUser(username string, sid int) error {
	d := dao.UserDao{}
	return d.AddStoreUser(username, sid)
}

func (u *UserService) UpdateMoney(username string, money float32) error {
	d := dao.UserDao{}
	return d.UpdateMoney(username, money)

}

//是否已存在该github账号的用户
func (u *UserService) IsExistGithubLogin(login string) (bool, error) {
	d := dao.UserDao{}
	_, err := d.SelectUserByGithubLogin(login)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

//通过电话查找用户
func (u *UserService) LoginBySms(phone string) (model.User, error) {
	d := dao.UserDao{}
	return d.SelectUserByPhone(phone)
}

func (u *UserService) JudgePhone(phone string) (bool, error) {
	d := dao.UserDao{}
	_, err := d.SelectUserByPhone(phone)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (u *UserService) SendCodeByPhone(phone string) (string, error) {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	err := SendMessage(phone, code)
	if err != nil {
		return "", err
	}
	return code, err
}

//验证码注册
func (u *UserService) RegisterBySms(userinfo model.User) error {
	d := dao.UserDao{}
	return d.RegisterBySms(userinfo)
}

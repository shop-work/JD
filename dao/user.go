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
	_, err := DB.Exec("insert into shop.userinfo(username, password,name,salt)values(?,?,?,?)", user.Username, user.Password, user.Name, user.Salt)
	return err
}

// SelectUserByUsername 查看用户详细信息
func (dao *UserDao) SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}
	row := DB.QueryRow("select uid,password,gender,name,phone,money,address_id,group_id,store_id,salt from shop.userinfo where username=?", username)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.Uid, &user.Password, &user.Gender, &user.Name, &user.Phone, &user.Money, &user.AddressId, &user.GroupId, &user.StoreId, &user.Salt)
	if err != nil {
		return user, err
	}
	user.Username = username
	return user, err
}

// SelectBasicUserByUsername 查看用户固定信息
func (dao *UserDao) SelectBasicUserByUsername(username string) (model.User, error) {
	user := model.User{}
	row := DB.QueryRow("select uid,password,group_id,store_id from shop.userinfo where username=?", username)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.Uid, &user.Password, &user.GroupId, &user.StoreId)
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

// UpdatePhone 更新用户电话
func (dao *UserDao) UpdatePhone(username string, phone string) error {
	_, err := DB.Exec("update shop.userinfo set phone=? where username=?", phone, username)
	return err
}

// UpdateName 更新昵称
func (dao *UserDao) UpdateName(username string, name string) error {
	_, err := DB.Exec("update shop.userinfo set name=? where username=?", name, username)
	return err
}

// UpdateGender 更新性别
func (dao *UserDao) UpdateGender(username string, gender string) error {
	_, err := DB.Exec("update shop.userinfo set gender=? where username=?", gender, username)
	return err
}

func (dao *UserDao) UpdateAddressId(username string, id int) error {
	_, err := DB.Exec("update shop.userinfo set address_id=? where address_id=?", username, id)
	return err
}

func (dao *UserDao) UpdateGroupId(username string, i int) error {
	_, err := DB.Exec("update shop.userinfo set group_id=? where username=?", i, username)
	return err
}

// AddStoreUser 商家入铺
func (dao *UserDao) AddStoreUser(username string, sid int) error {
	_, err := DB.Exec("update shop.userinfo set store_id=? where username=?", sid, username)
	return err
}

// UpdateMoney 更新余额
func (dao *UserDao) UpdateMoney(username string, money float32) error {
	_, err := DB.Exec("update shop.userinfo set money=? where username=?", money, username)
	return err
}

func (dao *UserDao) SelectUserByGithubLogin(login string) (model.User, error) {
	user := model.User{}
	row := DB.QueryRow("select uid,gender,name,phone,money,address_id,group_id,store_id,salt from shop.userinfo where username=?", login)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.Uid, &user.Gender, &user.Name, &user.Phone, &user.Money, &user.AddressId, &user.GroupId, &user.StoreId, &user.Salt)
	if err != nil {
		return user, err
	}
	user.GithubLogin = login
	return user, err
}

//通过电话查找用户
func (dao *UserDao) SelectUserByPhone(phone string) (model.User, error) {
	user := model.User{}
	row := DB.QueryRow("select uid,username,gender,name,money,address_id,group_id,store_id,salt from shop.userinfo where phone=?", phone)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.Uid, &user.Username, &user.Gender, &user.Name, &user.Money, &user.AddressId, &user.GroupId, &user.StoreId, &user.Salt)
	if err != nil {
		return user, err
	}
	user.Phone = phone
	return user, err
}

//验证码注册
func (dao *UserDao) RegisterBySms(user model.User) error {
	_, err := DB.Exec("insert into shop.userinfo(username, phone)values(?,?)", user.Username, user.Phone)
	return err
}

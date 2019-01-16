package model

import (
	"go_webtest/pkg/e"
	"time"
)

type User struct {//表模型
	Model
	Username string
	Mobile string
	Password string
	LastLogin time.Time
	LoginTimes int
}

var UserModels = User{}//产生一个对象

/*
验证电话和密码是否正确
 */

func(UserModel *User) UserLogin(mobile ,password string)(error string,user *User)  {

	db.First(&UserModel,"mobile = ?",mobile)
	if UserModel.ID == 0 {
		return  "账号不存在",nil
	}
	if e.MD5(password) != UserModel.Password{
		return  "密码不正确",nil
	}
	UserModel.LoginTimes++ //登录次数加一
	UserModel.LastLogin = time.Now() //登录时间改变
	db.Save(&UserModel)

	return "", UserModel
}

/*
验证session里的id和密码是否匹配
 */
func (UserModel *User)UserCheckSession(id,password string)(err string,user *User)  {

	db.First(&UserModel,"id = ? AND password = ?",id,password)
	if UserModel.ID == 0 {
		return "没有权限",nil
	}
	return "",UserModel
}

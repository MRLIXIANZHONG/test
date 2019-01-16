package admin

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_webtest/model"
	"strconv"
)

/*
登录
 */

func(Admin *admin) UserLogin(c *gin.Context)  {
	//接受参数
	mobile := c.PostForm("mobile")
	password :=c.PostForm("password")

	//查询是否有这条记录
	err,user :=model.UserModels.UserLogin(mobile,password)
	// 给前端返回状态
	if err !=""{
		ReturnJSON(c, &Result{200, err, nil})//返回前端错误信息
		return
	}
	//开启session
	session := sessions.Default(c)
	// 保存登录信息
	session.Set("user_session", strconv.Itoa(int(user.ID))+"|"+user.Password)
	_ = session.Save()
	fmt.Println("user_session = ", session.Get("user_session"))

	// 给前端返回状态
	ReturnJSON(c, &Result{200, err, nil})



}
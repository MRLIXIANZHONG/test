package admin

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_webtest/model"
	"go_webtest/pkg/e"
	"net/http"
	"strings"
)

type Result struct {
	Code int    `json:"code"` //返回状态码
	Msg  string `json:"msg"`  //返回提示信息
	Data *gin.H `json:"data"` //返回内容
}

type admin struct {
	//admin类
	User    *model.User
	Context *gin.Context
	Session sessions.Session
}

var Admin = admin{}

type Error struct {
	T   int		//跳转时间
	Url string	//跳转路径
	Msg string	//提示信息
}
var AdminError = Error{}

func ReturnJSON(c *gin.Context, result *Result) {
	if result.Code == 200 {
		result.Msg = "success"
	}
	c.JSON(200, result)

	c.Abort() //如果授权失败（例如密码不匹配），则调用abort以确保其余的处理程序不调用此请求。
}

func(Admin *admin)Checksession(c *gin.Context) {
	session := sessions.Default(c)
	//获取session
	user_session := session.Get("user_session")
	if user_session == nil{
		AdminError.T = 2
		AdminError.Url = "/"
		AdminError.Msg = "没有权限"
		c.Redirect(http.StatusFound, "/err")//重定向
		return
	}
	s := strings.Split(user_session.(string), "|")
	id := s[0]
	password := s[1]

	if id == "" || !e.IsNumeric(id){
		AdminError.T = 2
		AdminError.Url = "/"
		AdminError.Msg = "没有权限"
		c.Redirect(http.StatusFound, "/err")//重定向
		return
	}

	err,user := model.UserModels.UserCheckSession(id,password)
	// 给前端返回状态
	if err != "" {
		/*ReturnJSON(c, &Result{-1, err, nil})
		return*/
		AdminError.T = 2
		AdminError.Url = "/"
		AdminError.Msg = err
		c.Redirect(http.StatusFound, "/err")//重定向
		return
	}
	Admin.User = user

	c.Next()


}
/*
错误模板
 */
func (AdminError *Error) Errors(c *gin.Context) {
	//fmt.Printf("adminerror = %+v",AdminError)
	c.HTML(http.StatusOK, "layouts/err.html", gin.H{
		"msg":  AdminError.Msg,
		"url":  AdminError.Url,
		"time": AdminError.T,
	})
}


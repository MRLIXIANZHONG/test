package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_webtest/pkg/setting"
	"net/http"
	"go_webtest/controller/admin"
)

func InitRouter()*gin.Engine  {
	//设置中间键
	r := gin.New()//创建一个没有任何中间键的路由
	r.Use(gin.Logger())//日志写入中间键

	r.Use(gin.Recovery())//错误中恢复中间键

	gin.SetMode(setting.ServerSetting.RunMode)//获取调试模式

	// store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store := sessions.NewCookieStore([]byte("session.secret.password"))
	r.Use(sessions.Sessions("session", store))//开启加载session
	r.StaticFS("/static/",http.Dir("static"))//处理静态文件

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Header("Cache-Contorl", "public, max-age=2592000")//session存储时间2592000一个月时间
		c.File("./static/favicon.ico")//页面标题图标
	})

	r.LoadHTMLGlob("view/**/*")//设置模板路径
	r.GET("/", func(c *gin.Context) {
		c.HTML(200,"login/login.html",nil)
	})
	r.POST("/login",admin.Admin.UserLogin)

	//重定向错误页面
	r.GET("/err",admin.AdminError.Errors)

	admin1:=r.Group("/admin")
	admin1.Use(admin.Admin.Checksession)
	{
		admin1.GET("/", func(c *gin.Context) {
			c.HTML(200,"index/index.html",nil)
		})

	}


	return r




}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_webtest/model"
	"go_webtest/pkg/setting"
	"go_webtest/router"
	"net/http"
)

func main() {
	var(
		routers *gin.Engine
		s *http.Server
	)
	setting.Setup()//执行配置
	model.Setup()//链接数据库

	routers = router.InitRouter()
	s = &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),//地址端口
		Handler:        routers,		//路由头
		ReadTimeout:    setting.ServerSetting.ReadTimeout,//读取时间限制
		WriteTimeout:   setting.ServerSetting.WriteTimeout,//写入时间限制
		MaxHeaderBytes: 1 << 20,	//最大头字节
	}

	s.ListenAndServe()
}

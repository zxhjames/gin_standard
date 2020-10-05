package main

import (
	"Pro/gin_standard/global"
	"Pro/gin_standard/internal/routers"
	setting2 "Pro/gin_standard/pkg/setting"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

/**
* @program: src
* @description:
* @author: 占翔昊
* @create 2020-10-05 18:43
**/

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
		router := routers.NewRouter()
		s := &http.Server{
			Addr: ":"+global.ServerSetting.HttpPort,
			Handler: router,
			ReadTimeout: global.ServerSetting.ReadTimeout,
			WriteTimeout: global.ServerSetting.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServe()
}

func init() {
	if err := setupSetting();err!=nil {
		log.Error("init err:%v",err)
	}
}

func setupSetting() error {
	setting,err := setting2.NewSettings()
	if err != nil {
		return err
	}
	if err := setting.ReadSection("Server",&global.ServerSetting);err!=nil {
		return err
	}
	if err := setting.ReadSection("App",&global.AppSetting);err!=nil {
		return err
	}
	if err := setting.ReadSection("Database",&global.DatabaseSetting);err!=nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
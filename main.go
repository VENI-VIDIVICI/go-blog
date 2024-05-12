package main

import (
	"fmt"
	"github.com/VENI-VIDIVICI/go-blog/global"
	"github.com/VENI-VIDIVICI/go-blog/internal/model"
	"github.com/VENI-VIDIVICI/go-blog/internal/routers"
	"github.com/VENI-VIDIVICI/go-blog/pkg/logger"
	"github.com/VENI-VIDIVICI/go-blog/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := SetSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = SetDbEngine()
	if err != nil {
		log.Fatalf("init.SetDbEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}
func main() {
	router := routers.NewRouter()
	fmt.Println(global.AppSetting)
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
	s.ListenAndServe()
}

func SetSetting() error {
	settingInstance, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = settingInstance.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = settingInstance.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = settingInstance.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}

func SetDbEngine() error {
	//DbEngine
	db, err := model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.DbEngine = db
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

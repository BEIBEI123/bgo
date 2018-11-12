package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"

	"bgo/models"
	"bgo/pkg/config"
	"bgo/pkg/logging"
	"bgo/routers"
)

func main() {
	config.Setup()  //配置信息
	models.Setup()  //
	logging.Setup() //日志

	routersInit := routers.InitRouter()
	readTimeout := config.ServerSetting.ReadTimeout
	writeTimeout := config.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()

	log.Printf("Actual pid is %d", syscall.Getpid())

}

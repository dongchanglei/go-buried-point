package main

import (
	"gitlab.kanche.com/kkche/go-buried-point/web-gin/appsetting"
	"gitlab.kanche.com/kkche/go-buried-point/web-gin/db/mysql"
	"gitlab.kanche.com/kkche/go-buried-point/web-gin/routers"
	"fmt"
	"log"
	"syscall"
	"github.com/fvbock/endless"
	"runtime"
	"net/http"
)

func main() {

	appsetting.Setup()
	mysql.Setup()

	routersInit := routers.InitRouter()
	readTimeout := appsetting.ServerSetting.ReadTimeout
	writeTimeout := appsetting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", appsetting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	if runtime.GOOS == "windows" {
		server := &http.Server{
			Addr:           endPoint,
			Handler:        routersInit,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			MaxHeaderBytes: maxHeaderBytes,
		}

		server.ListenAndServe()
		return
	}

	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, routersInit)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

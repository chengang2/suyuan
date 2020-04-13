package main

import (
	"fmt"
	"log"
	"net/http"
	"suyuan/g"
	"suyuan/routers"
)

func main() {


	routersInit := routers.InitRouter()
	readTimeout := g.ServerSetting.ReadTimeout
	writeTimeout := g.ServerSetting.WriteTimeout

	endPoint := fmt.Sprintf(":%d", g.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	_ = server.ListenAndServe()

}

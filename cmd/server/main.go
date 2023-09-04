package main

import (
	webserver "github.com/WindBlog/module/http"
	"github.com/wonderivan/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := webserver.InitDefaultWebserver()
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	// 等待信号
	<-signalChan
	webserver.GetDefaultWebserverShutdownChan() <- syscall.SIGINT
}

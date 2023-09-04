package webserver

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"os"
	"syscall"
	"time"
)

var DefaultWebserverShutdownChan chan os.Signal
var DefaultWebserver *Webserver

const (
	DefaultAddress = "0.0.0.0:5000"
)

type Webserver struct {
	server       *http.Server
	shutdownChan chan os.Signal
}

func (w *Webserver) Init(address string) error {
	w.shutdownChan = make(chan os.Signal)
	r := gin.Default()
	// TODO logic route group addition
	//rg := r.Group("/")
	//rg.POST("/test", func(c *gin.Context) {
	//	logger.Info("test")
	//  ctx.Json({})
	//	return
	//})
	r.GET("/", func(c *gin.Context) {
		logger.Info("test /")
		c.JSON(200, "")
	})
	var err error

	go func() {
		w.server = &http.Server{Addr: address, Handler: r.Handler()}
		err = w.server.ListenAndServe()
	}()

	go func() {
		select {
		case signal := <-w.shutdownChan:
			switch signal {
			case syscall.SIGINT:
				logger.Info("webserver shutdown ...")
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if err = w.server.Shutdown(ctx); err != nil {
					logger.Fatal("shutdown %v", err)
				}
			}
		}
	}()

	if err != nil {
		logger.Fatal("webserver init: %v", err)
		return err
	}
	return err
}

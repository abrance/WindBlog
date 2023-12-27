package webserver

import "os"

// 使用说明
// InitDefaultWebserver 在只需要实现全局单例时使用
// 全局单例下调用	webserver.GetDefaultWebserverShutdownChan() <- syscall.SIGINT 则 webserver shutdown

// GetDefaultWebserverShutdownChan
// shutdown webserver
func GetDefaultWebserverShutdownChan() chan os.Signal {
	return DefaultWebserverShutdownChan
}

// InitDefaultWebserver
// 初始化默认的 webserver
func InitDefaultWebserver() error {
	DefaultWebserver = &Webserver{}
	err := DefaultWebserver.Init(DefaultAddress)
	DefaultWebserverShutdownChan = DefaultWebserver.shutdownChan
	return err
}

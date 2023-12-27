package log

import (
	"github.com/wonderivan/logger"
)

func Init() {
	err := logger.SetLogger("./log.json")
	if err != nil {
		return
	}
	logger.Debug("debug init : %v ", logger.LevelMap)
}

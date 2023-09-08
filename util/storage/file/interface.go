package file

import (
	"fmt"
	"github.com/wonderivan/logger"
	"os"
)

func GetRealPath(path string) string {
	return fmt.Sprintf("%s/%s", FileStoragePrefix)
}

func Init() {
	FileStoragePrefix, _ = os.Getwd()
	logger.Debug("file storage prefix : %s", FileStoragePrefix)
}

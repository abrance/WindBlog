package file

import (
	"fmt"
	"github.com/wonderivan/logger"
	"os"
)

func GetRealPath(path string) string {
	return fmt.Sprintf("%s/doc/%s", FileStoragePrefix, path)
}

func Init() {
	FileStoragePrefix, _ = os.Getwd()
	logger.Debug("file storage prefix : %s", FileStoragePrefix)
}

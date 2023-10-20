package storage

import (
	"github.com/WindBlog/util/storage/sqlite"
)

func Init() {
	// file.Init()
	// json_storage.Init()
	sqlite.Init()
}

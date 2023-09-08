package storage

import (
	"github.com/WindBlog/util/storage/file"
	"github.com/WindBlog/util/storage/json_storage"
)

func Init() {
	file.Init()
	json_storage.Init()
}

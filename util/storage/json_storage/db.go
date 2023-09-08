package json_storage

import (
	"github.com/WindBlog/util/errors"
	"github.com/dgraph-io/badger/v3"
	"github.com/wonderivan/logger"
	"os"
)

var (
	db        *badger.DB
	fileTable *FileTable
)

func Init() {
	var err error
	opt := badger.DefaultOptions("./db")
	db, err = badger.Open(opt)
	if err != nil {
		logger.Fatal(errors.JsonDBInitError)
		os.Exit(errors.JsonDBInitError)
	}

	// init file table
	fileTable = &FileTable{}
	fileTable.Init()
}

func GetFileTable() *FileTable {
	return fileTable
}

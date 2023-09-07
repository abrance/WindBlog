package json

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

func init() {
	var err error
	opt := badger.DefaultOptions("./db")
	db, err = badger.Open(opt)
	if err != nil {
		logger.Fatal(errors.JsonDBInitError)
		os.Exit(errors.JsonDBInitError)
	}
	defer db.Close()

	// init file table
	fileTable = &FileTable{}
	fileTable.SetTableName("file")
}

func GetFileTable() *FileTable {
	return fileTable
}

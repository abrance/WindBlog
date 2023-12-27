package sqlite

import (
	"github.com/WindBlog/util/errors"
	"github.com/WindBlog/util/storage/sqlite/table"
	"github.com/wonderivan/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var (
	db *gorm.DB
)

func Init() {
	initDB()
}

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price float64
}

func initDB() *gorm.DB {
	var err error
	db, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})

	if err != nil {
		logger.Fatal(errors.SqliteInitError)
		os.Exit(errors.SqliteInitError)
	}
	// Migrate the schema
	logger.Debug("start migrate tables ...")
	err = db.AutoMigrate(&table.Tag{}, &table.File{})

	if err != nil {
		logger.Fatal(errors.SqliteMigrateError)
		os.Exit(errors.SqliteMigrateError)
	}
	// db.Create(&Product{Name: "Mobile", Price: 500.50})
	return db
}

func GetDB() *gorm.DB {
	return db
}

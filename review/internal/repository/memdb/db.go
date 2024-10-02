package memdb

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDb(fileName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

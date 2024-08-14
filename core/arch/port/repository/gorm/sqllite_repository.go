package gorm

import (
	"sync"

	"gorm.io/driver/sqlite"
	realGorm "gorm.io/gorm"

	"river0825/cleanarchitecture/core/dependency/storage/gorm"
)

var (
	db  *realGorm.DB
	mu  sync.Mutex
	err error
)

func NewSqlLiteDB() *gorm.DB {
	if db == nil {
		mu.Lock()
		defer mu.Unlock()
		if db == nil {
			db, err = realGorm.Open(sqlite.Open("file::memory:?cache=shared"), &realGorm.Config{
				IgnoreRelationshipsWhenMigrating: true,
			})
			if err != nil {
				return nil
			}
		}
	}

	return &gorm.DB{DB: db}
}

func AutoMigrateSqlLiteDB(models []interface{}) error {
	return db.AutoMigrate(models...)
}

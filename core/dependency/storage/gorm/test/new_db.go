package test

import (
	"fmt"
	"log"
	"os"
	"time"

	gorm2 "gorm.io/gorm"
	"gorm.io/gorm/logger"

	"river0825/cleanarchitecture/core/dependency/storage/gorm"
)

func NewTestDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		"localhost", 5430, "postgres", "postgres", "postgres", "dev")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	return gorm.NewDB(dsn, &gorm2.Config{
		Logger: newLogger,
	})
}

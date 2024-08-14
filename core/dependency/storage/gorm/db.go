package gorm

import (
	"context"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"river0825/cleanarchitecture/core/dependency/gin_engine/checkalive"
)

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Taipei"

// NewDB opens a database specified by its database driver name and a driver-specific data source name, usually consisting of at least a database name and connection information.
// @param dsn string
// @return *gorm.DB
func NewDB(dsn string, config ...*gorm.Config) *DB {
	c := &gorm.Config{}
	if len(config) > 0 {
		c = config[0]
	}

	db, err := gorm.Open(sqlite.Open("gorm.db"), c)

	db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Session(&gorm.Session{SkipDefaultTransaction: true})
	})

	if err != nil {
		panic(err)
	}

	return &DB{
		DB: db,
	}
}

func NewDBWithLogger(dsn string, config ...*gorm.Config) *DB {
	newLogger := logger.New(
		log.New(os.Stdout, "", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	return NewDB(dsn, append(config, &gorm.Config{Logger: newLogger})...)
}

var _ checkalive.ICheckAlive = (*DB)(nil)

type DB struct {
	DB *gorm.DB
}

func (db *DB) Conn(ctx context.Context) *gorm.DB {
	return db.DB.WithContext(ctx)
}

func (db *DB) Ping(ctx context.Context) error {
	r := db.DB.WithContext(ctx).Exec("SELECT 1")
	return r.Error
}

func (db *DB) Name() string {
	return "gorm/postgres"
}

func (db *DB) Close() {
	// do nothing
}

func (db *DB) FilterPublishedAt(gormDB *gorm.DB) *gorm.DB {
	gormDB.Where("published_at is not null")
	return gormDB
}

func (db *DB) PublishedAtNotNull(tx *gorm.DB, tableName string) *gorm.DB {
	return tx.Where(tableName + ".published_at is not null")
}

func (db *DB) GormDB() *gorm.DB {
	return db.DB
}

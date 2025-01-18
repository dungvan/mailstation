package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func GetDBClient() *gorm.DB {
	return db
}

func New(dbURL string, dbDebugMode bool) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: func(dbDebugMode bool) logger.Interface {
			if dbDebugMode {
				return logger.Default.LogMode(logger.Info)
			}
			return logger.Default.LogMode(logger.Silent)
		}(dbDebugMode),
	})
	if err != nil {
		return nil, fmt.Errorf("Open Connection %w", err)
	}

	sql, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("SQL %w", err)
	}

	sql.SetMaxIdleConns(20)
	sql.SetConnMaxLifetime(time.Hour)
	sql.SetMaxOpenConns(30)

	if err = sql.Ping(); err != nil {
		return nil, fmt.Errorf("Ping Error %w", err)
	}

	return db, nil
}

func Close(db *gorm.DB) {
	if db != nil {
		sql, err := db.DB()
		if err != nil {
			sql.Close()
		}
	}
}

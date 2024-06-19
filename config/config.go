package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlLite: %v", err)

	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *logger {
	//Initialize logger
	logger = NewLogger(p)
	return logger
}

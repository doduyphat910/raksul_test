package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	onceDB      sync.Once
	singletonDB *gorm.DB
)

func InitDBConnection(dsn string, config gorm.Config) (*gorm.DB, error) {
	var err error
	onceDB.Do(func() {
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}), &config)
		if err == nil {
			singletonDB = db
		}
	})
	if err != nil {
		return nil, err
	}
	return singletonDB, nil
}

func GetDB() *gorm.DB {
	return singletonDB
}

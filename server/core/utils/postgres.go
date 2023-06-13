package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB *gorm.DB
)

type PgExecutor struct {
	DB *gorm.DB
}

func InitDB(url string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitPool(url string, maxConnections int) (*gorm.DB, error) {
	db, err := InitDB(url)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxOpenConns(maxConnections)
	db.LogMode(true)

	DB = db
	return DB, nil
}

func (p *PgExecutor) GetDB() *gorm.DB {
	return p.DB
}

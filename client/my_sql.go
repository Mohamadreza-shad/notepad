package client

import (
	"github.com/Mohamadreza-shad/notepad/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DbConnection()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

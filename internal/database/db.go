package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB  *gorm.DB
	log logrus.FieldLogger
}

func NewDatabase(log logrus.FieldLogger) (Database, error) {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		return Database{}, err
	}

	return Database{
		DB:  db,
		log: log,
	}, nil
}

func (db *Database) StartMigrations() error {
	return db.DB.AutoMigrate(&Image{})
}

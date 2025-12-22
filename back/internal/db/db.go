package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Fatal("Error connecting to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.WithError(err).Fatal("Error initialization database")
	}
	if err := sqlDB.Ping(); err != nil {
		logrus.WithError(err).Fatal("Not access database")
	}

	if err := db.AutoMigrate(); err != nil {
		logrus.WithError(err).Fatal("Error migrating database")
	}

	logrus.Info("Successfully connected to database")
	return db
}

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}

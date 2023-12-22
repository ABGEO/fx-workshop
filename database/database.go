package database

import (
	"fmt"

	"github.com/abgeo/fx-workshop/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(conf *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(conf.Database.FilePath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return db, nil
}

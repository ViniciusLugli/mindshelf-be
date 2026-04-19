package repositories

import (
	"errors"

	"github.com/ViniciusLugli/mindshelf/internal/utils/envutil"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := envutil.DatabaseDSN()

	if dsn == "" {
		return nil, errors.New("DSN or DATABASE_URL is not configured")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

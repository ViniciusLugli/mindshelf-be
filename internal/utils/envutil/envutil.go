package envutil

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnvIfPresent() error {
	err := godotenv.Load()
	if err == nil {
		return nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	var pathErr *os.PathError
	if errors.As(err, &pathErr) && errors.Is(pathErr.Err, os.ErrNotExist) {
		return nil
	}

	return err
}

func DatabaseDSN() string {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("DSN")
	}

	return dsn
}

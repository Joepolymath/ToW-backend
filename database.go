package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	UserCredentials     string
	PasswordCredentials string
	Host                string
	Database            string
	FallbackHost        string
	FallbackDatabase    string
	FallbackUser        string
	FallbackPassword    string
	// New in 1.1.1
	Logger logger.Interface
}

func PrepareDialectConnection(dialect, dsn string) (gorm.Dialector, error) {
	sqlDB, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "prepare dialect connection failed")
	}
	return postgres.New(postgres.Config{Conn: sqlDB}), nil
}

func GetDBConnection(config *DBConfig) (db *gorm.DB, err error) {
	var (
		connectionType,
		dialect,
		dsn string
	)

	if len(config.UserCredentials) > 0 && len(config.PasswordCredentials) > 0 {
		dialect = "postgres"
		connectionType = "Local"
		dsn = fmt.Sprintf("sslmode=disable host=%s dbname=%s user=%s password=%s",
			config.Host,
			config.Database,
			config.UserCredentials,
			config.PasswordCredentials)

		log.Printf("Using dialect `%s` with connection type `%s`\n", dialect, connectionType)
		dialector, err := PrepareDialectConnection(dialect, dsn)
		if err != nil {
			return nil, err
		}

		return gorm.Open(dialector, &gorm.Config{Logger: config.Logger})
	}
	return nil, fmt.Errorf("Error connecting to db")
}
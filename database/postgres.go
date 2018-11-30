package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mirzakhany/pkg/database/postgresql"
	"github.com/mirzakhany/pkg/logger"
)

// InitPostgres init postgres
func InitPostgres(maxTry int, settings PostgresSettings) (*gorm.DB, error) {

	var err error
	var database *gorm.DB

	configs := postgresql.ConnectionURL{
		User:     settings.Username,
		Password: settings.Password,
		Host:     settings.Host,
		Socket:   settings.Socket,
		Database: settings.DatabaseName,
		Options:  settings.Options,
	}

	for {
		database, err = gorm.Open("postgres", configs.String())
		if err == nil {
			break
		}
		logger.LogError.Errorf("Connect to Postgres failed du error: %v", err)
		if maxTry > 0 {
			maxTry--
		} else {
			return nil, err
		}
	}
	logger.LogAccess.Info("Connection to Postgres stabilised")
	return database, err
}

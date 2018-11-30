package database

import (
	"github.com/jinzhu/gorm"
	"github.com/mirzakhany/pkg/logger"
	"github.com/pkg/errors"
	"gopkg.in/gormigrate.v1"
)

var (
	// DB database instance
	DB *gorm.DB
)

// KV key value type
type KV map[string]string

// DBSettings settings of database
type DBSettings struct {
	Engine       string           `json:"engine"`
	MaxTry       int              `json:"max_try"`
	MySQL        MySQLSettings    `json:"mysql"`
	Postgres     PostgresSettings `json:"postgres"`
	SQLite       SQLiteSettings   `json:"sqlite"`
	DBModels     []interface{}
	DBMigrations []*gormigrate.Migration
	ForeignKeys  []ForeignKey
}

// MySQLSettings is settings of mySQL
type MySQLSettings struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	DatabaseName string `json:"database_name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Socket       string `json:"socket"`
	DialTimeout  int    `json:"dial_timeout"`
	Options      KV     `json:"options"`
	Timeout      int    `json:"timeout"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}

// PostgresSettings is settings of Postgres
type PostgresSettings struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	DatabaseName string `json:"database_name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Socket       string `json:"socket"`
	DialTimeout  int    `json:"dial_timeout"`
	Options      KV     `json:"options"`
}

// SQLiteSettings is settings of SQLite
type SQLiteSettings struct {
	DatabaseName string `json:"database_name"`
}

// ForeignKey struct of table foreign keys
type ForeignKey struct {
	Model       interface{}
	Field       string
	Destination string
	onDelete    string
	onUpdate    string
}

// InitDatabase for initialize database
func InitDatabase(settings *DBSettings) error {
	logger.LogAccess.Debugf("Init Database Engine as %s", settings.Engine)
	var err error
	switch settings.Engine {
	case "postgres":
		DB, err = InitPostgres(settings.MaxTry, settings.Postgres)
	case "sqlite":
		DB, err = InitSqlite(settings.MaxTry, settings.SQLite)
	case "mysql":
		DB, err = InitMySQLDB(settings.MaxTry, settings.MySQL)
	default:
		logger.LogError.Error("database error: can't find database driver")
		return errors.New("can't find database driver")
	}

	initSchema(DB, settings)

	return err
}

func initSchema(db *gorm.DB, settings *DBSettings) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, settings.DBMigrations)
	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			settings.DBModels...,
		)
		if err != nil {
			logger.LogError.Fatalf("auto migration failed, %v", err)
			return errors.New("auto migration failed")
		}

		for _, fk := range settings.ForeignKeys {
			if err := tx.Model(fk.Model).AddForeignKey(fk.Field, fk.Destination, fk.onDelete, fk.onUpdate).Error; err != nil {
				logger.LogError.Fatalf("add foreign-key failed, %v", err)
				return err
			}
		}
		return nil
	})
}

// CloseDatabase close database session
func CloseDatabase() {
	err := DB.Close()
	if err != nil {
		logger.LogError.Errorf("database error: can't close database, %v", err)
	}
}

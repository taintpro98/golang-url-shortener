package database

import (
	"fmt"
	"golang-url-shortener/config"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func PostgresqlDatabaseProvider() (*gorm.DB, error) {
	db, err := NewPostgresqlDatabase()
	if err != nil {
		log.Error().Err(err).Msg("connect to database error")
		return nil, err
	}
	return db, nil
}

func NewPostgresqlDatabase() (*gorm.DB, error) {
	dbConfig := config.DatabaseConfig{
		Host:         os.Getenv("POSTGRES_HOST"),
		DatabaseName: os.Getenv("POSTGRES_DB"),
		Username:     os.Getenv("POSTGRES_USER"),
		Password:     os.Getenv("POSTGRES_PASSWORD"),
		Port:         os.Getenv("POSTGRES_PORT"),
	}
	dsn := GetDatabaseDSN(dbConfig)
	customLogger := NewCustomLogger(dbLoggerConfig{
		ignoreRecordNotFoundError: false,
	})
	customLogger.logLevel = logger.Info
	client, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		},
	), &gorm.Config{
		Logger: customLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "public.",
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	_, err = client.DB()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetDatabaseDSN(DBConf config.DatabaseConfig) string {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s TimeZone=%s",
		DBConf.Host, DBConf.Port, DBConf.Username, DBConf.DatabaseName, "UTC",
	)

	if DBConf.SSLMode != "" {
		dsn += fmt.Sprintf(" sslmode=%s", DBConf.SSLMode)
	}

	if DBConf.Password != "" {
		dsn += fmt.Sprintf(" password=%s", DBConf.Password)
	}
	return dsn
}

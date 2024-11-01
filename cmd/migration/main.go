package main

import (
	"context"
	"flag"
	"fmt"
	"golang-url-shortener/config"
	"golang-url-shortener/pkg/database"
	"os"

	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("No .env file")
	}

	dir := flag.String("dir", "migrations", "Path to migrations")
	flag.Parse()
	args := flag.Args()
	command := args[0]

	ctx := context.Background()
	dbConfig := config.DatabaseConfig{
		Host:         os.Getenv("POSTGRES_HOST"),
		DatabaseName: os.Getenv("POSTGRES_DB"),
		Username:     os.Getenv("POSTGRES_USER"),
		Password:     os.Getenv("POSTGRES_PASSWORD"),
		Port:         os.Getenv("POSTGRES_PORT"),
	}
	dsn := database.GetDatabaseDSN(dbConfig)
	db, err := goose.OpenDBWithDriver("postgres", dsn)
	if err != nil {
		log.Warn().Err(err).Msg("open db error")
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Warn().Err(err).Msg("goose: failed to close DB")
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.RunContext(ctx, command, db, *dir, arguments...); err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("goose %v: %v", command, err))
	}
}

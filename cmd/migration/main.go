package main

import (
	"context"
	"flag"
	"golang-url-shortener/config"
	"golang-url-shortener/pkg/database"
	"os"

	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
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
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.RunContext(ctx, command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}

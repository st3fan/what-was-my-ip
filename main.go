package main

import (
	"context"
	_ "embed"
	"log"
	"net/url"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/tech10/ipify"
)

//go:embed setup.sql
var setupSQL string

func databaseURLFromEnv() (string, error) {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")),
		Host:   os.Getenv("DB_HOSTNAME"),
		Path:   os.Getenv("DB_DATABASE"),
	}
	return u.String(), nil
}

func setup() {
	databaseURL, err := databaseURLFromEnv()
	if err != nil {
		log.Fatalf("Failed to get database url: %s", err)
	}

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	if _, err := conn.Exec(context.Background(), setupSQL); err != nil {
		log.Fatalf("Failed to setup database: %s", err)
	}
}

func lookup() {
	address, err := ipify.GetIp4()
	if err != nil {
		log.Fatalf("failed to get ip: %s", err)
	}

	databaseURL, err := databaseURLFromEnv()
	if err != nil {
		log.Fatalf("Failed to get database url: %s", err)
	}

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}

	defer conn.Close(context.Background())

	if _, err := conn.Exec(context.Background(), "CALL InsertAddressIfChanged($1)", address); err != nil {
		log.Fatalf("Failed to insert into database: %s", err)
	}
}

func main() {
	switch os.Args[1] {
	case "setup":
		setup()
	case "lookup":
		lookup()
	}
}

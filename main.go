package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/tech10/ipify"
)

//go:embed setup.sql
var setupSQL string

func setup() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	if _, err := conn.Exec(context.Background(), setupSQL); err != nil {
		log.Fatalf("failed to setup database: %s", err)
	}
}

func lookup() {
	address, err := ipify.GetIp4()
	if err != nil {
		log.Fatalf("failed to get ip: %s", err)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
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

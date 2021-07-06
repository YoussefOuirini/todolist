package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"

	_ "github.com/lib/pq"
)

const dbUrl = "postgres://postgres:example@postgres:5432/%s?sslmode=disable"

func main() {
	if len(os.Args) < 2 {
		panic("Required argument action")
	}

	cmd := os.Args[1]

	migrateService(cmd, "todo")
	migrateService(cmd, fmt.Sprintf("%s-test", "todo"))
}

func migrateService(cmd string, dbname string) {
	// make sure db exists
	postgresDb := mustOpenDB(fmt.Sprintf(dbUrl, "postgres"))

	row := postgresDb.QueryRow(fmt.Sprintf(
		"SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s');",
		dbname,
	))

	var dbExists bool

	if err := row.Scan(&dbExists); err != nil {
		panic(err)
	}

	if !dbExists {
		fmt.Printf("create database: %s\n", dbname)

		if _, err := postgresDb.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", dbname)); err != nil {
			panic(err)
		}
	}

	// connect to the database
	db := mustOpenDB(fmt.Sprintf(dbUrl, dbname))

	// run migrations
	sourcePath := "cmd/db/migrations/"

	fmt.Printf("[%s] running migrations %s\n", dbname, cmd)

	switch cmd {
	case "up":
		if err := goose.Up(db.DB, sourcePath); err != nil {
			fmt.Printf("[%s] error: %s\n", dbname, err.Error())
		}
	case "down":
		if err := goose.Down(db.DB, sourcePath); err != nil {
			fmt.Printf("[%s] error: %s\n", dbname, err.Error())
		}
	}
}

func mustOpenDB(connString string) *sqlx.DB {
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
		log.Fatalf("failed to connect to db")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("failed to ping db")
	}

	return db
}

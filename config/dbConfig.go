package config

import (
	"database/sql"
	"net/url"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func InitDB(databaseUrl string) *sql.DB {
	parseDBUrl, _ := url.Parse(databaseUrl)
	db, err := sql.Open(parseDBUrl.Scheme, databaseUrl)
	if err != nil {
		panic(err)
	}

	runMigration, errRM := strconv.ParseBool(LoadEnv("ACTIVATE_DB_MIGRATION"))
	if errRM != nil {
		return db
	}

	if runMigration {
		RunDBMigration(db)
	}

	return db
}

func RunDBMigration(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:./migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}

	m.Up()
}

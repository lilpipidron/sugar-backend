package postgresql

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

type Storage struct {
	db *sql.DB
}

func New(psqlInfo, dbname string) (*Storage, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	migration, err := migrate.NewWithDatabaseInstance("file:/migration", dbname, driver)
	if err != nil {
		return nil, err
	}

	if err := migration.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		}
	}

	return &Storage{db: db}, nil
}

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
  const errFunc = "storage.postgresql.NewStorage"

  db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil,  fmt.Errorf("%s: %w", errFunc, err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil,  fmt.Errorf("%s: %w", errFunc, err)
	}

	migration, err := migrate.NewWithDatabaseInstance("file:/migration", dbname, driver)
	if err != nil {
		return nil,  fmt.Errorf("%s: %w", errFunc, err)
	}

	if err := migration.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
      return nil,  fmt.Errorf("%s: %w", errFunc, err)
		}
	}

	return &Storage{db: db}, nil
}

//TODO querys

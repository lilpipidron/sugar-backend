package user

import (
	"database/sql"
	"time"

	"github.com/lilpipidron/sugar-backend/internal/models/users"
)

type Repository interface {
	AddNewUser(user users.User) error
	FindUser(login, password string) (users.User, error)
	DeleteUserById(userID int64) error
	ChangeName(login, newName string) error
	ChangeBirhday(login string, newBirthday time.Time) error
	ChangeGender(login string, newGender users.Gender) error
	ChangeWeight(login string, newWeight int)
	ChangeCarbohydrateRatio(login string, newCarbohydrateRatio int) error
	ChangeBreadUnit(login string, newBreadUnit int) error
}
type repository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *repository {
	return &repository{DB: db}
}

func (db *repository) AddNewUser(user users.User) error {
	return nil
}

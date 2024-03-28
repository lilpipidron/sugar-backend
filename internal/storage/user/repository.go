package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lilpipidron/sugar-backend/internal/models/users"
)

type Repository interface {
	AddNewUser(user users.User, password string) error
	FindUser(login, password string) (users.User, error)
	DeleteUserById(userID int64) error
	ChangeName(login, newName string) error
	ChangeBirhday(login string, newBirthday time.Time) error
	ChangeGender(login string, newGender users.Gender) error
	ChangeWeight(login string, newWeight int) eroro
	ChangeCarbohydrateRatio(login string, newCarbohydrateRatio int) error
	ChangeBreadUnit(login string, newBreadUnit int) error
}
type repository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *repository {
	return &repository{DB: db}
}

func (db *repository) AddNewUser(user users.User, password string) error {
	const op = "storage.user.AddNewUser"
	query := "INSERT INTO users (login, password) VALUES ($1, $2) RETURNING user_id"
	id, err := db.DB.Exec(query, user.Login, password)
	if err != nil {
		return fmt.Errorf("%s: failer add new user: %w", op, err)
	}

	query = "INSERT INTO user_info (userID, name, birthday, gender, weight, carbohydrateRatio, breadUnit) VALUSE ($1, $2, $3, $4, $5, $6, $7)"

	return nil
}

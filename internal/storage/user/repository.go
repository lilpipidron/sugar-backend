package user

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lilpipidron/sugar-backend/internal/models/users"
)

type Repository interface {
	AddNewUser(user users.User, password string) error
	FindUser(login, password string) (*users.User, error)
	DeleteUser(userID int64) error
	ChangeName(userID int64, newName string) error
	ChangeBirthday(userID int64, newBirthday time.Time) error
	ChangeGender(userID int64, newGender users.Gender) error
	ChangeWeight(userID int64, newWeight int) error
	ChangeCarbohydrateRatio(userID int64, newCarbohydrateRatio int) error
	ChangeBreadUnit(userID int64, newBreadUnit int) error
	ChangeAllInfo(userID int64, newName string, newBirthday time.Time, newGender users.Gender, newWeight int, newCarbohydrate int, newBreadUnit int, newHeight int) error
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
	_, err := db.DB.Exec(query, user.Login, password)

	if err != nil {
		return fmt.Errorf("%s: failed add new user: %w", op, err)
	}

	var id int64
	err = db.DB.QueryRow("SELECT lastval()").Scan(&id)

	if err != nil {
		return fmt.Errorf("%s: failed get last user id: %w", op, err)
	}

	query = "INSERT INTO user_info (user_id, name, birthday, gender, weight, carbohydrate_ratio, bread_unit, height) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = db.DB.Exec(query, id, user.UserInfo.Name, user.UserInfo.Birthday, user.UserInfo.Gender, user.UserInfo.Weight, user.UserInfo.CarbohydrateRatio, user.UserInfo.BreadUnit, user.UserInfo.Height)
	if err != nil {
		return fmt.Errorf("%s: failed add user info: %w", op, err)
	}

	return nil
}

func (db *repository) FindUser(login, password string) (*users.User, error) {
	const op = "storage.user.FindUser"

	query := "SELECT * FROM users where login = $1 and password = $2"
	row, err := db.DB.Query(query, login, password)
	if err != nil {
		return nil, fmt.Errorf("%s: failed find user: %w", op, err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(fmt.Errorf("%s: failed close user's row: %w", op, err))
		}
	}(row)

	u := &users.User{}
	row.Next()
	err = row.Scan(&u.UserID, &u.Login, &u.Password)
	if err != nil {
		return nil, fmt.Errorf("%s: failed scan row: %w", op, err)
	}

	query = "SELECT * FROM user_info where user_id = $1"
	row, err = db.DB.Query(query, u.UserID)
	if err != nil {
		return nil, fmt.Errorf("%s: failed find user_info: %w", op, err)
	}
	row.Next()
	err = row.Scan(&u.UserID, &u.UserInfo.Name, &u.UserInfo.Birthday, &u.UserInfo.Gender, &u.UserInfo.Weight, &u.UserInfo.CarbohydrateRatio, &u.UserInfo.BreadUnit, &u.UserInfo.Height)
	if err != nil {
		return nil, fmt.Errorf("%s: failed scan user_info: %w", op, err)
	}
	return u, nil
}

func (db *repository) DeleteUser(userID int64) error {
	const op = "storage.user.DeleteUser"

	query := "DELETE FROM user_info WHERE user_id = $1"
	_, err := db.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("%s: failed delete user_info: %w", op, err)
	}

	query = "DELETE FROM users WHERE user_id = $1"
	_, err = db.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("%s: failed delete user: %w", op, err)
	}

	return nil
}

func (db *repository) ChangeName(userID int64, newName string) error {
	const op = "storage.user.ChangeName"

	query := "UPDATE user_info SET name = $2 WHERE user_id = $1"
	_, err := db.DB.Exec(query, userID, newName)
	if err != nil {
		return fmt.Errorf("%s: failed change name: %w", op, err)
	}

	return nil
}

func (db *repository) ChangeBirthday(userID int64, newBirthday time.Time) error {
	const op = "storage.user.ChangeBirthday"

	query := "UPDATE user_info SET birthday = $2 WHERE user_id = $1"
	_, err := db.DB.Exec(query, userID, newBirthday)
	if err != nil {
		return fmt.Errorf("%s: failed change birhday: %w", op, err)
	}

	return nil
}

func (db *repository) ChangeGender(userID int64, newGender users.Gender) error {
	const op = "storage.user.ChangeGender"

	query := "UPDATE user_info SET gender = $2 WHERE user_id = $1"
	_, err := db.DB.Exec(query, userID, newGender)
	if err != nil {
		return fmt.Errorf("%s: failed change gender: %w", op, err)
	}

	return nil
}

func (db *repository) ChangeWeight(userID int64, newWeight int) error {
	const op = "storage.user.ChangeWeight"

	query := "UPDATE user_info SET weight = $2 WHERE user_id = $1"
	_, err := db.DB.Exec(query, userID, newWeight)
	if err != nil {
		return fmt.Errorf("%s: failed change weight: %w", op, err)
	}

	return nil
}

func (db *repository) ChangeCarbohydrateRatio(userID int64, newCarbohydrateRation int) error {
	const op = "storage.user.ChangeCarbohydrateRatio"

	query := "UPDATE user_info SET carbohydrate_ratio = $2 WHERE user_id = $1"
	_, err := db.DB.Exec(query, userID, newCarbohydrateRation)
	if err != nil {
		return fmt.Errorf("%s: failed change carbohydrate ratio: %w", op, err)
	}

	return nil
}

func (db *repository) ChangeBreadUnit(userID int64, newBreadUnit int) error {
	const op = "storage.user.ChangeBreadUnit"

	query := "UPDATE user_info SET bread_unit = $2 WHERE user_id = $1"
	_, err := db.DB.Exec(query, userID, newBreadUnit)
	if err != nil {
		return fmt.Errorf("%s: failed change bread unit: %w", op, err)
	}

	return nil
}

func (db *repository) ChangeAllInfo(userID int64, newName string, newBirthday time.Time, newGender users.Gender, newWeight int, newCarbohydrate int, newBreadUnit int, newHeight int) error {
	const op = "storage.user.ChangeAllInfo"

	err := db.ChangeName(userID, newName)
	if err != nil {
		return err
	}

	err = db.ChangeBirthday(userID, newBirthday)
	if err != nil {
		return err
	}

	err = db.ChangeGender(userID, newGender)
	if err != nil {
		return err
	}

	err = db.ChangeWeight(userID, newWeight)
	if err != nil {
		return err
	}

	err = db.ChangeCarbohydrateRatio(userID, newCarbohydrate)
	if err != nil {
		return err
	}

	err = db.ChangeBreadUnit(userID, newBreadUnit)
	if err != nil {
		return err
	}

	query := "UPDATE user_info SET height = $2 WHERE user_id = $1"
	_, err = db.DB.Exec(query, userID, newHeight)
	if err != nil {
		return fmt.Errorf("%s: failed change height: %w", op, err)
	}

	return nil
}

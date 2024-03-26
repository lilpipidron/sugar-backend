package storage

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrNoteNotFound    = errors.New("note not found")
	ErrProductNotFound = errors.New("product not found")

	ErrInvalidName              = errors.New("invalid name")
	ErrInvalidDate              = errors.New("invalid date")
	ErrInvalidWeight            = errors.New("invalid weight")
	ErrInvalidSugarLevel        = errors.New("invalid sugar level")
	ErrInvalidCarbsValue        = errors.New("invalid carbs value")
	ErrInvalidBreadUnitsValue   = errors.New("invalid bread units value")
	ErrInvalidCarbohydrateRatio = errors.New("invalid carbohydrate ratio")

	ErrUserExists    = errors.New("user already exists")
	ErrNoteExists    = errors.New("note already exists")
	ErrProductExists = errors.New("product already exists")
)

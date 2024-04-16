package request

import (
	"time"

	"github.com/lilpipidron/sugar-backend/internal/models/users"
)

 type AddUser struct {
	Login             string       `json:"login"`
	Password          string       `json:"password"`
	Name              string       `json:"name"`
	Birthday          time.Time    `json:"birthday"`
	Gender            users.Gender `json:"gender"`
  Weight            int          `json:"weight"`
  CarbohydrateRatio int          `json:"carbohydrate-ratio"`
	BreadUnit         int          `json:"bread-unit"`
}

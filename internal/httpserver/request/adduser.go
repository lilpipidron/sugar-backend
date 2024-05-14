package request

import (
	"encoding/json"
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

func (u *AddUser) UnmarshalJSON(data []byte) error {
	type Alias AddUser
	aux := &struct {
		Birthday string `json:"birthday"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	u.Birthday, err = time.Parse("2006-01-02", aux.Birthday)
	return err
}

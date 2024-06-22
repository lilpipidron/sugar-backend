package request

import (
	"encoding/json"
	"github.com/lilpipidron/sugar-backend/internal/models/users"
	"time"
)

type ChangeUserInfo struct {
	ID              int64        `json:"id"`
	NewName         string       `json:"new_name"`
	NewBirthday     time.Time    `json:"new_birthday"`
	NewBreadUnit    int          `json:"new_bread_unit"`
	NewCarbohydrate int          `json:"new_carbohydrate_ratio"`
	NewGender       users.Gender `json:"new_gender"`
	NewWeight       int          `json:"new_weight"`
	NewHeight       int          `json:"new_height"`
}

func (u *ChangeUserInfo) UnmarshalJSON(data []byte) error {
	type Alias ChangeUserInfo
	aux := &struct {
		NewBirthday string `json:"new_birthday"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	u.NewBirthday, err = time.Parse("2006-01-02", aux.NewBirthday)
	return err
}

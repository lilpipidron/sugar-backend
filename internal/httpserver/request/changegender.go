package request

import "github.com/lilpipidron/sugar-backend/internal/models/users"

type ChangeGender struct {
	ID        int64        `json:"id"`
	NewGender users.Gender `json:"new_gender"`
}

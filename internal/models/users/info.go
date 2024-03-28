package users

import "time"

type UserInfo struct {
	Name              string    `json:"name"`
	Birhday           time.Time `json:"birthday"`
	Gender            Gender    `json:"gender"`
	Weight            int       `json:"weight"`
	CarbohydrateRatio int       `json:"carbohydrate-ratio"`
	BreadUnit         int       `json:"bread-unit"`
}

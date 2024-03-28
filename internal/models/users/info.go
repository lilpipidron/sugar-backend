package users

type UserInfo struct {
	Name              string `json:"name"`
	Gender            Gender `json:"gender"`
	Weight            int    `json:"weight"`
	CarbohydrateRatio int    `json:"carbohydrate-ratio"`
	GrainUnit         int    `json:"grain-unit"`
}

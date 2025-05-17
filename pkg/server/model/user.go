package model

type User = struct {
	UserID  string `json:"userId"`
	IconUrl string `json:"iconUrl"`
	Cd      int    `json:"cd"`
	Power   int    `json:"power"`
	Weight  int    `json:"weight"`
	Volume  int    `json:"volume"`
	Point   []int  `json:"point"`
}

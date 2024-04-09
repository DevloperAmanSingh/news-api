package models

type ApiUser struct {
	Id        string `json:"id"`
	Karma     int    `json:"karma"`
	Submitted []int  `json:"submitted"`
	Comments  []int  `json:"comments"`
	
}

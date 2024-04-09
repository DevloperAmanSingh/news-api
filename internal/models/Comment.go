package models

type Comment struct {
	By      string    `json:"by"`
	ID      int       `json:"id"`
	Kids    []int     `json:"kids"`
	Parent  int       `json:"parent"`
	Text    string    `json:"text"`
	Time    int       `json:"time"`
	Type    string    `json:"type"`
	Replies []Comment `json:"replies,omitempty"`
}

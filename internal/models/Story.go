package models

type Story struct {
	ID          int       `json:"id"`
	By          string    `json:"by"`
	Descendants int       `json:"descendants"`
	Kids        []int     `json:"kids"`
	Score       int       `json:"score"`
	Time        int       `json:"time"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	URL         string    `json:"url"`
	Comments    []Comment `json:"comments,omitempty"`
}

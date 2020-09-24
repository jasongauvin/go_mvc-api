package models

type Activity struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

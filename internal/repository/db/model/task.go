package model

type Task struct {
	ID     int64  `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Status int    `json:"status" db:"status"`
}

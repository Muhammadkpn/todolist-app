package model

type Task struct {
	ID     int64  `gorm:"column:id;primary_key"`
	Title  string `gorm:"column:title"`
	Status int    `gorm:"column:status"`

	// TableName string `gorm:"TableName:task"` // Replace with your actual table name
}

func (t *Task) TableName() string {
	return "task" // Explicitly specify the table name
}

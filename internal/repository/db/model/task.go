package model

import "time"

type Task struct {
	ID          int       `gorm:"column:id;primaryKey;not null"`
	Title       string    `gorm:"column:title;not null"`
	Description string    `gorm:"column:description"`
	UserID      int       `gorm:"column:user_id"`
	Label       []Label   `gorm:"many2many:task_label;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status      string    `gorm:"column:status;not null"`
	DueDate     time.Time `gorm:"column:due_date"`
	CreatedBy   string    `gorm:"column:created_by;not null"`
	CreatedDate time.Time `gorm:"column:created_date"`
	UpdatedBy   string    `gorm:"column:updated_by"`
	UpdatedDate time.Time `gorm:"column:updated_date"`
}

func (t *Task) TableName() string {
	return "task"
}

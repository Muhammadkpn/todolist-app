package model

import "time"

type Label struct {
	ID          int       `gorm:"column:id;primaryKey;not null"`
	Name        string    `gorm:"column:name;not null"`
	UserID      int       `gorm:"column:user_id"`
	CreatedBy   string    `gorm:"column:created_by;not null"`
	CreatedDate time.Time `gorm:"column:created_date"`
	UpdatedBy   string    `gorm:"column:updated_by"`
	UpdatedDate time.Time `gorm:"column:updated_date"`
}

func (t *Label) TableName() string {
	return "label"
}

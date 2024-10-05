package model

import "time"

type User struct {
	ID          int       `gorm:"column:id;primaryKey;not null"`
	Username    string    `gorm:"column:username;not null;unique"`
	Email       string    `gorm:"column:email;not null;unique"`
	Password    string    `gorm:"column:password;not null"`
	CreatedBy   string    `gorm:"column:created_by;not null"`
	CreatedDate time.Time `gorm:"column:created_date"`
	UpdatedBy   string    `gorm:"column:updated_by"`
	UpdatedDate time.Time `gorm:"column:updated_date"`
}

func (t *User) TableName() string {
	return "user"
}

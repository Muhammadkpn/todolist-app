package model

import "base/pkg/enums/petType"

type Pet struct {
	ID      uint64 `gorm:"primary_key"`
	Name    string
	PetType petType.Type
}

func (Pet) TableName() string {
	return "pet"
}

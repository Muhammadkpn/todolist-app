package model

type TaskLabel struct {
	TaskID  int `gorm:"primaryKey;foreignKey:TaskID;references:ID"`
	LabelID int `gorm:"primaryKey;foreignKey:LabelID;references:ID"`
}

func (t *TaskLabel) TableName() string {
	return "task_label"
}

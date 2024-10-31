package task

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Status   int
	Type     int
	UserID   uint // 发起者ID，0表示系统发起
	Progress int
	Error    string `gorm:"type:text"`
	Props    string `gorm:"type:text"`
}

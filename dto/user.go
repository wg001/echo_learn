package dto

import (
	"time"
)

type UserORM struct {
	Id         uint      `json:"id" gorm:"primary_key"`
	UserName   string    `gorm:"column:username" json:"user_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Level      int64     `json:"level"`
	IsDelete   uint8     `gorm:"column:is_delete"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type User interface {
	GetAllUser() []UserORM
	GetOneUser(condition string) UserORM
}

package dto

import "time"

type user struct {
	Id         uint
	UserName   string    `gorm:"user_name"`
	Email      string    `gorm:"email"`
	Phone      string    `gorm:"phone"`
	Address    string    `gorm:"address"`
	City       string    `gorm:"city"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
}

type User interface {
	GetAllUser() []user
	GetOneUser(condition string) user
}

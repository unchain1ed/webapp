package entity

import (
	"time"
)

type User struct {
	ID uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt" gorm:"primarykey"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"primarykey"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"index"`
	UserId string `json:"userId" binding:"required,min=2,max=10"`
	Password string `json:"password" binding:"required,min=4,max=20"`
}

type UserIdChange struct {
	NowId string `json:"nowId" binding:"required,min=2,max=10"`
	ChangeId string `json:"changeId" binding:"required,min=2,max=10"`
}

type UserPwChange struct {
	UserId string `json:"userId" binding:"required,min=2,max=10"`
	NowPassword string `json:"nowPassword" binding:"required,min=4,max=20"`
	ChangePassword string `json:"changePassword" binding:"required,min=4,max=20"`
}

type FormUser struct {
	UserId string `json:"userId" binding:"required,min=2,max=10"`
	Password string `json:"password" binding:"required,min=4,max=20"`
}
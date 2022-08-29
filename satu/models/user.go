package models

import (
	"time"
)

type User struct {
	Id        uint      `gorm:"type:bigint(20);not null" json:"id" form:"id"`
	Name      string    `gorm:"type:varchar(45)" json:"name" form:"name"`
	UserName  string    `gorm:"type:varchar(45)" json:"user_name" form:"user_name"`
	Password  string    `gorm:"type:varchar(225)" json:"password" form:"password"`
	CreateAt  string    `json:"created_at" form:"created_at"`
	CreateBy  time.Time `gorm:"type:bigint(20)" json:"created_by" form:"created_by"`
	UpdatedAt string    `json:"updated_at" form:"updated_at"`
	UpdatedBy time.Time `gorm:"type:bigint(20)" json:"updated_by" form:"updated_by"`
}

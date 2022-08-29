package models

import (
	"time"
)

type Merchant struct {
	Id           uint      `gorm:"type:bigint(20);not null" json:"id" form:"id"`
	UserId       uint      `gorm:"type:int(40);not null" json:"user_id" form:"user_id"`
	MerchantName string    `gorm:"type:varchar(40);not null" json:"merchant_name" form:"merchant_name"`
	CreateAt     string    `json:"created_at" form:"created_at"`
	CreateBy     time.Time `gorm:"type:bigint(20)" json:"created_by" form:"created_by"`
	UpdatedAt    string    `json:"updated_at" form:"updated_at"`
	UpdatedBy    time.Time `gorm:"type:bigint(20)" json:"updated_by" form:"updated_by"`
}

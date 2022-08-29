package models

import (
	"time"
)

type Merchant struct {
	Id           uint      `gorm:"type:bigint(20);not null" json:"id" form:"id"`
	UserId       uint      `gorm:"type:int(40);not null" json:"user_id" form:"user_id"`
	User         User      `gorm:"foreignKey:UserId" json:"user" form:"user"`
	MerchantName string    `gorm:"type:varchar(40);not null" json:"merchant_name" form:"merchant_name"`
	CreatedAt    time.Time `gorm:"type:bigint(20)" json:"created_at" form:"created_at"`
	CreatedBy    uint      `json:"created_by" form:"created_by"`
	UpdatedAt    time.Time `gorm:"type:bigint(20)" json:"updated_at" form:"updated_at"`
	UpdatedBy    uint      `json:"updated_by" form:"updated_by"`
}

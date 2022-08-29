package models

import (
	"time"
)

type Outlet struct {
	Id         uint      `gorm:"type:bigint(20);not null" json:"id" form:"id"`
	MerchantId uint      `gorm:"type:bigint(20);not null" json:"merchant_id" form:"merchant_id"`
	Merchant   Merchant  `gorm:"foreignKey:MerchantId" json:"merchant" form:"merchant"`
	OutletName string    `gorm:"type:varchar(40);not null" json:"outlet_name" form:"outlet_name"`
	CreatedAt  time.Time `gorm:"type:bigint(20)" json:"created_at" form:"created_at"`
	CreatedBy  uint      `json:"created_by" form:"created_by"`
	UpdatedAt  time.Time `gorm:"type:bigint(20)" json:"updated_at" form:"updated_at"`
	UpdatedBy  uint      `json:"updated_by" form:"updated_by"`
}

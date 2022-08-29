package models

import (
	"time"
)

type Transaction struct {
	Id         uint      `gorm:"type:bigint(20);not null" json:"id" form:"id"`
	MerchantId uint      `gorm:"type:bigint(20);not null" json:"merchant_id" form:"merchant_id"`
	OutletId   uint      `gorm:"type:bigint(20);not null" json:"outlet_id" form:"outlet_id"`
	BillTotal  uint      `gorm:"type:double;not null" json:"bill_total" form:"bill_total"`
	CreateAt   string    `json:"created_at" form:"created_at"`
	CreateBy   time.Time `gorm:"type:bigint(20)" json:"created_by" form:"created_by"`
	UpdatedAt  string    `json:"updated_at" form:"updated_at"`
	UpdatedBy  time.Time `gorm:"type:bigint(20)" json:"updated_by" form:"updated_by"`
}

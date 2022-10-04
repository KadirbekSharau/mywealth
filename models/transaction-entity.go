package models

import (
	"gorm.io/gorm"
)

// EntityTransactions is database entity for each transaction
type EntityTransactions struct {
	gorm.Model
	Name        string `gorm:"type:varchar(30)"`
	Description string `gorm:"type:varchar"`
	Amount      int    `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	AccountID   uint   `gorm:"not null"`
	CategoryID  uint   `gorm:"not null"`
}

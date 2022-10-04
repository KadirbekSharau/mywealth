package models

import (
	"gorm.io/gorm"
)

// EntityAccounts is database entity for finance acccount
type EntityAccounts struct {
	gorm.Model
	Name                string `gorm:"type:varchar(30)"`
	Active              bool   `gorm:"not null"`
	UserID              uint   `gorm:"not null"`
	AccountTransactions []EntityTransactions `gorm:"foreignKey:AccountID"`
}

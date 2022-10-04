package models

import (

	//"github.com/KadirbekSharau/bookswap-backend/util"
	"gorm.io/gorm"
)

// EntityCategories is database entity for categories
type EntityCategories struct {
	gorm.Model
	Name                 string `gorm:"type:varchar(30)"`
	UserID               uint   `gorm:"not null"`
	CategoryTransactions []EntityTransactions `gorm:"foreignKey:CategoryID"`
}

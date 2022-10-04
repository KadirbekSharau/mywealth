package models

import (
	"time"

	"github.com/KadirbekSharau/mywealth-backend/util"
	"gorm.io/gorm"
)

// EntityUsers is database entity for user
type EntityUsers struct {
	gorm.Model
	Name             string               `gorm:"type:varchar(30)"`
	Surname          string               `gorm:"type:varchar(30)"`
	AvatarUrl        string               `gorm:"type:varchar"`
	Email            string               `gorm:"type:varchar(50);unique;not null"`
	Password         string               `gorm:"type:varchar(255)"`
	IsAdmin          bool                 `gorm:"type:bool;default:false"`
	Active           bool                 `gorm:"type:bool"`
	UserAccounts     []EntityAccounts     `gorm:"foreignKey:UserID"`
	UserCategories   []EntityCategories   `gorm:"foreignKey:UserID"`
	UserTransactions []EntityTransactions `gorm:"foreignKey:UserID"`
}

func (entity *EntityUsers) BeforeCreate(db *gorm.DB) error {
	entity.Password = util.HashPassword(entity.Password)
	return nil
}

func (entity *EntityUsers) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

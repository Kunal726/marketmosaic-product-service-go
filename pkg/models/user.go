package models

import (
	"time"
)

// User represents the user entity in the database
type User struct {
	Id           uint      `gorm:"primaryKey;column:id"`
	Email        string    `gorm:"column:email"`
	Name         string    `gorm:"column:name"`
	Roles        string    `gorm:"column:roles"`
	BusinessName string    `gorm:"column:business_name"`
	ContactPhone string    `gorm:"column:contact_phone"`
	Address      string    `gorm:"column:address"`
	DateAdded    time.Time `gorm:"column:date_added"`
	IsActive     bool      `gorm:"column:is_active"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "marketmosaic_users"
}

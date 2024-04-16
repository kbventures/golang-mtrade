package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Name     string    `gorm:"type:text;not null"`
	Username string    `gorm:"uniqueIndex;not null"`
	Password string    `gorm:"type:text;not null"`
	// Assuming `emailVerified` is stored as a timestamp without time zone in the database
	EmailVerified string `gorm:"type:timestamp(3) without time zone"`
	Image         string `gorm:"type:text"` // Assuming `image` is stored as text in the database
}

// TableName sets the table name for the User model
func (User) TableName() string {
	return "user"
}

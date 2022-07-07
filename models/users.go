package models

import (
	"github/Cyantosh0/demo/lib"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        lib.BinaryUUID `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Address   string         `json:"address"`
	Age       int            `json:"age"`
	Gender    bool           `json:"gender"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// BeforeCreate run this before creating user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = lib.BinaryUUID(id)
	return err
}

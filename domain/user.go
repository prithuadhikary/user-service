package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primarykey"`
	Username string
	Password string
	Role     string
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New()

	// Digest and store the hex representation.
	digest := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(digest[:])

	return nil
}

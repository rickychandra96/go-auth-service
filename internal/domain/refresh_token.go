package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	Token     string         `gorm:"uniqueIndex;not null" json:"token"`
	ExpiresAt time.Time      `gorm:"not null" json:"expires_at"`
	IsRevoked bool           `gorm:"default:false" json:"is_revoked"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (rt *RefreshToken) BeforeCreate(tx *gorm.DB) error {
	if rt.ID == uuid.Nil {
		rt.ID = uuid.New()
	}
	return nil
}

func (*RefreshToken) TableName() string {
	return "refresh_tokens"
}

// IsExpired Helper method untuk cek apakah token sudah expired
func (rt *RefreshToken) IsExpired() bool {
	return time.Now().After(rt.ExpiresAt)
}

// IsValid Helper method untuk cek apakah token valid (tidak expired dan tidak revoked)
func (rt *RefreshToken) IsValid() bool {
	return !rt.IsExpired() && !rt.IsRevoked
}

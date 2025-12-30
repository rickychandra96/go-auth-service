package domain

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        uuid.UUID  `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	User      User       `json:"-"`
	Token     string     `json:"token"`
	ExpiresAt time.Time  `json:"expires_at"`
	IsRevoked bool       `json:"is_revoked"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

func (rt *RefreshToken) BeforeCreate() {
	if rt.ID == uuid.Nil {
		rt.ID = uuid.New()
	}
	rt.CreatedAt = time.Now()
	rt.UpdatedAt = time.Now()
	rt.IsRevoked = false
}

// IsExpired Helper method untuk cek apakah token sudah expired
func (rt *RefreshToken) IsExpired() bool {
	return time.Now().After(rt.ExpiresAt)
}

// IsValid Helper method untuk cek apakah token valid (tidak expired dan tidak revoked)
func (rt *RefreshToken) IsValid() bool {
	return !rt.IsExpired() && !rt.IsRevoked
}

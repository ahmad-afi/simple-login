package user

import (
	"time"
)

type UserEntity struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Role      string    `json:"role" db:"role"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

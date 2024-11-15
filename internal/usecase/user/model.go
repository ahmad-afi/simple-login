package user

type CreateUserReq struct {
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

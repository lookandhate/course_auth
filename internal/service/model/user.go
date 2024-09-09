package model

import "time"

type UserRole int

const (
	UserUnknownRole UserRole = iota
	Admin
	User
)

// CreateUserModel user representation for creation at service level.
type CreateUserModel struct {
	Name            string   `json:"name,omitempty"`
	Email           string   `json:"email,omitempty"`
	Password        string   `json:"password,omitempty"`
	PasswordConfirm string   `json:"password_confirm,omitempty"`
	Role            UserRole `json:"role,omitempty"`
}

// UserModel at service level.
type UserModel struct {
	ID        int
	Name      string
	Email     string
	Role      int
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UpdateUserModel user update data on service level.
type UpdateUserModel struct {
	ID       int
	Name     *string
	Email    *string
	Password *string
	Role     int
}

// CreateUserRepositoryModel user info for creation at repository level.
type CreateUserRepositoryModel struct {
	Name     string
	Email    string
	Password string
	Role     UserRole
}

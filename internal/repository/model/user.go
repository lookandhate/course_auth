package repoModel

import "time"

type UserRole int

type CreateUserModel struct {
	Name     string
	Email    string
	Password string
	Role     UserRole
}

// UserModel at repo level.
type UserModel struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Role      int       `db:"role"`
	Password  string    `db:"password_hash"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

package repository

type UserRole int

type CreateUserModel struct {
	Name     string
	Email    string
	Password string
	Role     UserRole
}

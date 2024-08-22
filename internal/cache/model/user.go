package model

type UserModel struct {
	ID          int64  `redis:"id"`
	Name        string `redis:"name"`
	Email       string `redis:"email"`
	Role        int    `redis:"role"`
	Password    string `redis:"password_hash"`
	CreatedAtNS int64  `redis:"created_at_ns"`
	UpdatedAtNS int64  `redis:"updated_at_ns"`
}

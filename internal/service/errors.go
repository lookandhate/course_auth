package service

import "errors"

var (
	// ErrInvalidID - when passed an invalid ID (<= 0 for integer IDs for example).
	ErrInvalidID = errors.New("invalid id")

	// ErrInvalidRole - when passed invalid user role on user create or update.
	ErrInvalidRole = errors.New("invalid role")

	// ErrPasswordMismatch - when password and password confirmations are not the same.
	ErrPasswordMismatch = errors.New("password mismatch")

	// ErrUserDoesNotExist - when trying to access not existing user.
	ErrUserDoesNotExist = errors.New("user does not exist")

	// ErrEmptyUser - when passed empty user.
	ErrEmptyUser = errors.New("empty user")
)

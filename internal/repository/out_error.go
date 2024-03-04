package repository

import "errors"

const (
	errMsgUserNotFound = "user not found"
	errMsgEmailIsTaken = "user with this email already exists"
)

var (
	ErrUserNotFound = errors.New(errMsgUserNotFound) // ErrUserNotFound ...
	ErrEmailIsTaken = errors.New(errMsgEmailIsTaken) // ErrEmailIsTaken ...
)

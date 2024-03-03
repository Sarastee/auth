package model

import "time"

type User struct {
	ID        int64
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCreate struct {
	Name     string
	Email    string
	Password string
	Role     string
}

type UserUpdate struct {
	ID    int64
	Name  *string
	Email *string
	Role  *string
}

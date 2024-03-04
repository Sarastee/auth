package model

import "time"

// User ...
type User struct {
	ID        int64
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserCreate ...
type UserCreate struct {
	Name     string
	Email    string
	Password string
	Role     string
}

// UserUpdate ...
type UserUpdate struct {
	ID    int64
	Name  *string
	Email *string
	Role  *string
}

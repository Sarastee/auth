package model

import "time"

// UserRole ...
type UserRole string

const (
	UNKNOWN UserRole = "UNKNOWN" // UNKNOWN ...
	ADMIN   UserRole = "ADMIN"   // ADMIN ...
	USER    UserRole = "USER"    // USER ...
)

// User ...
type User struct {
	ID        int64
	Name      string
	Email     string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserCreate ...
type UserCreate struct {
	Name     string
	Email    string
	Password string
	Role     UserRole
}

// UserUpdate ...
type UserUpdate struct {
	ID    int64
	Name  *string
	Email *string
	Role  *UserRole
}

package models

import (
	"guardian-api/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Repository database acessor methods
type Repository interface {
	FindByID() *User
	FindByEmail() *User
	Save() error
}

// User schema
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	DarkMode  bool      `db:"dark_mode" json:"dark_mode"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// Authenticate returns true if the password matches
func (u *User) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// EncryptPassword encrypts the user password and returns an nil or error
func (u *User) EncryptPassword() error {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(encrypted)
	return nil
}

// Valid checks if the user is valid
func (u *User) Valid() bool {
	return len(u.Name) > 0 &&
		utils.ValidateEmail(u.Email) &&
		utils.ValidatePassword(u.Password)
}

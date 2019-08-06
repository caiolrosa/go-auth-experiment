package user

import (
	"fmt"
	"guardian-api/db"
	"guardian-api/utils"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Repository database acessor methods
type Repository interface {
	FindByEmail(email string) *User
	Save() error
}

// User schema
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

// FindByEmail implementes the user repository method to find a user by email
func (u *User) FindByEmail() (User, error) {
	var userFound = User{}
	db, err := db.GetConnection()
	if err != nil {
		return userFound, err
	}
	defer db.Close()

	err = db.Where("email = ?", u.Email).Find(&userFound).Error
	if err != nil {
		return userFound, err
	}

	return userFound, nil
}

// Save implementes the user repository Create method to persist a user to the db
func (u *User) Save() error {
	db, err := db.GetConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Create(u).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
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

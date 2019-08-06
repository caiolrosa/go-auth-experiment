package user

import (
	"fmt"
	"guardian-api/db"
	"guardian-api/utils"

	"github.com/jinzhu/gorm"
)

// Repository database acessor methods
type Repository interface {
	FindAll() []User
	Save() error
}

// User schema
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

// FindAll implements the user repository FindAll method to find all users in db
func (u *User) FindAll() ([]User, error) {
	db, err := db.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
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

// Valid checks if the user is valid
func (u *User) Valid() bool {
	return len(u.Name) > 0 &&
		utils.ValidateEmail(u.Email) &&
		utils.ValidatePassword(u.Password)
}

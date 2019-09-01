package services

import (
	"guardian-api/db"
	"guardian-api/models"

	"github.com/labstack/gommon/log"
)

// IUserRepository exposes user database accessor methods
type IUserRepository interface {
	FindByID(ID int64) (models.User, error)
	FindByEmail(email string) (models.User, error)
	Save(user models.User) (models.User, error)
}

// UserRepository IUserRepository
type UserRepository struct {
	DBClient *db.Client
}

// FindByID implementes the user repository method to find a user by id
func (u *UserRepository) FindByID(ID int64) (models.User, error) {
	var userFound = models.User{}
	dbRef, err := u.DBClient.GetConnection()
	if err != nil {
		log.Error(err)
		return userFound, err
	}
	defer dbRef.Close()

	err = dbRef.Get(&userFound, "SELECT * FROM Users where id = ?", ID)
	if err != nil {
		return userFound, err
	}

	return userFound, nil
}

// FindByEmail implementes the user repository method to find a user by email
func (u *UserRepository) FindByEmail(email string) (models.User, error) {
	var userFound = models.User{}
	dbRef, err := u.DBClient.GetConnection()
	if err != nil {
		log.Error(err)
		return userFound, err
	}
	defer dbRef.Close()

	err = dbRef.Get(&userFound, "SELECT * FROM Users where email = ?", email)
	if err != nil {
		return userFound, err
	}

	return userFound, nil
}

// Save implementes the user repository Create method to persist a user to the db
func (u *UserRepository) Save(user models.User) (models.User, error) {
	newUser := models.User{}

	dbRef, err := u.DBClient.GetConnection()
	if err != nil {
		log.Error(err)
		return newUser, err
	}
	defer dbRef.Close()

	result, err := dbRef.NamedExec(
		"INSERT INTO Users (name, email, password) VALUES (:name, :email, :password)", user)
	if err != nil {
		return newUser, err
	}

	uid, err := result.LastInsertId()
	if err != nil {
		return newUser, err
	}

	newUser, err = u.FindByID(uid)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

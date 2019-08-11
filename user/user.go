package user

import (
	"guardian-api/db"
	"guardian-api/utils"
	"time"

	log "github.com/sirupsen/logrus"
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
	DBClient  db.API    `db:"-" json:"-"`
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// FindByID implementes the user repository method to find a user by id
func (u *User) FindByID() (User, error) {
	var userFound = User{}
	dbRef, err := u.DBClient.GetConnection()
	if err != nil {
		log.Error(err)
		return userFound, err
	}
	defer dbRef.Close()

	err = dbRef.Get(&userFound, "SELECT * FROM Users where id = ?", u.ID)
	if err != nil {
		return userFound, err
	}

	return userFound, nil
}

// FindByEmail implementes the user repository method to find a user by email
func (u *User) FindByEmail() (User, error) {
	var userFound = User{}
	dbRef, err := u.DBClient.GetConnection()
	if err != nil {
		log.Error(err)
		return userFound, err
	}
	defer dbRef.Close()

	err = dbRef.Get(&userFound, "SELECT * FROM Users where email = ?", u.Email)
	if err != nil {
		return userFound, err
	}

	return userFound, nil
}

// Save implementes the user repository Create method to persist a user to the db
func (u *User) Save() error {
	dbRef, err := u.DBClient.GetConnection()
	if err != nil {
		log.Error(err)
		return err
	}
	defer dbRef.Close()

	result, err := dbRef.NamedExec(
		"INSERT INTO Users (name, email, password) VALUES (:name, :email, :password)",
		u,
	)
	if err != nil {
		return err
	}

	uid, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = uid
	*u, err = u.FindByID()
	if err != nil {
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

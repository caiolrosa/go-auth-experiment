package migrations

import (
	"fmt"
	"guardian-api/db"
	"guardian-api/user"
)

// InitialMigration creates the base User schema
func initialMigration() {
	db, err := db.GetConnection()
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&user.User{})
}

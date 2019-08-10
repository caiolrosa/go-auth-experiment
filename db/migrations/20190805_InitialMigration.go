package migrations

import (
	"fmt"
	"guardian-api/db"
	"guardian-api/user"

	log "github.com/sirupsen/logrus"
)

// InitialMigration creates the base User schema
func initialMigration(dbClient db.API) {
	log.Info("Applying 20190805_InitialMigration")
	db, err := dbClient.GetConnection()
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&user.User{})
}

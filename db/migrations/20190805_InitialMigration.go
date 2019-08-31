package migrations

import (
	"guardian-api/db"

	log "github.com/sirupsen/logrus"
)

// initialMigration creates the base User schema
func initialMigration(dbClient db.API) {
	log.Info("Applying 20190805_InitialMigration")
	db, err := dbClient.GetConnection()
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	db.MustExec(`
		CREATE TABLE IF NOT EXISTS Users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
}

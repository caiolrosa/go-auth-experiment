package migrations

import (
	"guardian-api/db"

	log "github.com/sirupsen/logrus"
)

// Migrate runs all migrations
func Migrate(dbClient db.API) {
	log.Info("Starting to apply migrations...")
	initialMigration(dbClient)
	log.Info("Finished applying migrations.")
}

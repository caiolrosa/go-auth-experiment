package migrations

import log "github.com/sirupsen/logrus"

// Migrate runs all migrations
func Migrate() {
	log.Info("Starting to apply migrations...")
	initialMigration()
	log.Info("Finished applying migrations.")
}

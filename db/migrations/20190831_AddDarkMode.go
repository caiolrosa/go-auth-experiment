package migrations

import (
	"guardian-api/db"

	log "github.com/sirupsen/logrus"
)

// addDarkModeMigration adds dark mode setting to the user table
func addDarkModeMigration(dbClient db.API) {
	log.Info("Applying 20190830_AddDarkMode")
	db, err := dbClient.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Query(`SELECT * FROM pragma_table_info('Users') WHERE name = 'dark_mode'`)
	if err != nil {
		panic(err)
	}

	if result.Next() {
		log.Info("Skipping... Column already exists.")
		return
	}

	db.MustExec("ALTER TABLE Users ADD COLUMN dark_mode INTEGER DEFAULT 0")
}

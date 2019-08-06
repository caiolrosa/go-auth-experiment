package db

import (
	"fmt"
	"guardian-api/utils"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	gormDialect    = "sqlite3"
	gormDBPath     = "/db/data/hsvr_guardian.db"
	gormTestDBPAth = "/db/data/hsvr_guardian_test.db"
)

// GetConnection gets a connection to the database
func GetConnection() (*gorm.DB, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	env := utils.GetEnv()
	if env == utils.TestEnv {
		return gorm.Open(gormDialect, currentPath+gormTestDBPAth)
	}

	return gorm.Open(gormDialect, currentPath+gormDBPath)
}

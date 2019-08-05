package db

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	gormDialect = "sqlite3"
	gormDBPath  = "/db/data/hsvr_guardian.db"
)

// GetConnection gets a connection to the database
func GetConnection() (*gorm.DB, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return gorm.Open(gormDialect, currentPath + gormDBPath)
}

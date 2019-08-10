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
	gormTestDBPath = "/db/data/hsvr_guardian_test.db"
)

// API exposes the database api
type API interface {
	GetConnection() (*gorm.DB, error)
}

// Client implements the DBApi interface
type Client struct{}

// GetConnection gets a connection to the database
func (c *Client) GetConnection() (*gorm.DB, error) {
	return gorm.Open(gormDialect, getDBPath())
}

func getDBPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	env := utils.GetEnv()
	if env == utils.TestEnv {
		return currentPath+gormTestDBPath
	}

	return currentPath+gormDBPath
}

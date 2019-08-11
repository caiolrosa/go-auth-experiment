package db

import (
	"fmt"
	"guardian-api/utils"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	driverName = "sqlite3"
	dbPath     = "/db/data/hsvr_guardian.db"
	dbTestPath = "/db/data/hsvr_guardian_test.db"
)

// API exposes the database api
type API interface {
	GetConnection() (*sqlx.DB, error)
}

// Client implements the DBApi interface
type Client struct{}

// GetConnection gets a connection to the database
func (c *Client) GetConnection() (*sqlx.DB, error) {
	return sqlx.Open(driverName, getDBPath())
}

func getDBPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	env := utils.GetEnv()
	if env == utils.TestEnv {
		return currentPath + dbTestPath
	}

	return currentPath + dbPath
}

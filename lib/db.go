package lib

import (
	"database/sql"
	"errors"
	"os"
	"sync"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var (
	dbClient *sql.DB
	once     sync.Once
)

func GetDBUrl() (string, error) {
	dbUrl, exists := os.LookupEnv("DB_URL")
	if !exists {
		return "", errors.New("DB_URL environment variable not set")
	}
	return dbUrl, nil
}

func GetDBClient() (*sql.DB, error) {
	dbUrl, err := GetDBUrl()
	if err != nil {
		return nil, err
	}

	var dbErr error
	once.Do(func() {
		dbClient, dbErr = sql.Open("libsql", dbUrl)
	})

	return dbClient, dbErr
}

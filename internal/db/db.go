package database

import (
	database_tables "Pastely/internal/db/tables"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init(appName string, dbName string) error {

	dbPath := getDBPath(appName, dbName)
	createDBPath(dbPath)

	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the database is accessible
	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	database_tables.InitFoldersTable(DB)
	database_tables.InitItemsTable(DB)
	database_tables.InitTagsTable(DB)
	database_tables.InitTagAssignmentsTable(DB)

	return nil
}

func getDBPath(appName string, dbName string) string {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	// Ensure dbName ends with ".db"
	if !strings.HasSuffix(dbName, ".db") {
		dbName += ".db"
	}

	dbPath := filepath.Join(dir, appName, dbName)

	return dbPath
}

func createDBPath(dbPath string) error {
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		log.Fatalf("Failed to create directory for database: %v", err)
	}
	return nil
}

package store

import (
	"log"
	"os"
	"path/filepath"

	"database/sql"

	"github.com/fatih/color"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // Pulls in sqlite specific driver
)

var schema = `
	DROP TABLE IF EXISTS todo;
	CREATE TABLE todo (
		id INTEGER PRIMARY KEY,
		name VARCHAR(255) DEFAULT 0,
		completed VARCHAR(5) DEFAULT '',
		created_at DATETIME DEFAULT '',
		updated_at DATETIME DEFAULT '',
		completed_at DATETIME DEFAULT NULL
	);
`

// DataStore is the "base" for handling the database connection
type DataStore struct {
	ConnectionString string
	DB               *sqlx.DB
}

// Select wraps the DB.Select to make sure we have a valid connection pointer
func (ds *DataStore) Select(dest interface{}, query string, args ...interface{}) error {
	err := ds.Connect()
	if err != nil {
		return err
	}
	return ds.DB.Select(dest, query, args...)
}

// Get wraps the DB.Get to make sure we have a valid connection pointer
func (ds *DataStore) Get(dest interface{}, query string, args ...interface{}) error {
	err := ds.Connect()
	if err != nil {
		return err
	}
	return ds.DB.Get(dest, query, args...)
}

// MustExec wraps the DB.Select to make sure we have a valid connection pointer
func (ds *DataStore) MustExec(query string, args ...interface{}) sql.Result {
	err := ds.Connect()
	if err != nil {
		panic(err)
	}
	return ds.DB.MustExec(query, args...)
}

// Exec wraps the DB.Select to make sure we have a valid connection pointer
func (ds *DataStore) Exec(query string, args ...interface{}) (sql.Result, error) {
	err := ds.Connect()
	if err != nil {
		return nil, err
	}
	return ds.DB.Exec(query, args...)
}

// Connect connects to the database
func (ds *DataStore) Connect() error {
	if ds.DB != nil {
		return nil
	}
	if len(ds.ConnectionString) == 0 {
		log.Println("Defaulting the connection string...")
		ds.ConnectionString = "_dist/database/todos.db"
	}
	var err error
	dir, _ := filepath.Split(ds.ConnectionString)
	if dir != "" {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			log.Println("Creating directory and database...")
			os.MkdirAll(dir, 0777)
		}
	}
	ds.DB, err = sqlx.Connect("sqlite3", ds.ConnectionString)
	return err
}

// Init drops and recreates the database
func (ds *DataStore) Init() {
	log.Println("Initializing database...")
	color.Yellow("Warning: Dropping and re-creating tables...")
	ds.MustExec(schema)
}

package pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// InitDB Function to initiate database
func InitDB() (*sql.DB, error) {
	// DB sources
	const (
		host     = "db"
		port     = 5432
		user     = "postgres"
		password = "password"
		db_name  = "todo"
	)

	// Build the data source for connection fmt.Sprintf returns string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname = %s sslmode=disable",
		host, port, user, password, db_name,
	)

	// Open a connection to database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Check the database connection
	if err := db.Ping(); err != nil {
		// Close connection and return nil if Ping fails
		db.Close()
		return nil, err
	}

	// Run migration
	// RunMigrations(db)

	fmt.Println("Connected to database as 🚀")

	return db, nil
}

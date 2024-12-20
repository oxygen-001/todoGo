package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Function to run all database migrations
func RunMigrations(db *sql.DB) {
	// Path to the migrations folder
	migrationsDir := "migrations"

	// Read the migration file
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Error reading migration directory: %v", err)
	}

	// Loop through all migration files
	for _, file := range files {
		// Check if the file has .sql extension
		if filepath.Ext(file.Name()) == ".sql" {
			// That returns full path
			migrationFile := filepath.Join(migrationsDir, file.Name())
			// Print migration name
			fmt.Printf("Running migration: %s/n", file.Name())

			// Read the contents of the migration file
			migrationSQL, err := os.ReadFile(migrationFile)
			if err != nil {
				log.Fatalf("Error reading migration file %s: %v", migrationFile, err)
			}

			// Execute the SQL commands in the migration file on database
			_, err = db.Exec(string(migrationSQL))
			if err != nil {
				log.Fatalf("Error executing migration %s: %v", migrationFile, err)
			}

			// Print success message
			fmt.Printf("Migration %s completed successfully. \n", file.Name())
		}
	}
}

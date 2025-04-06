package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"reflect"

	"github.com/devndam/go-starter/app/models"
	"github.com/devndam/go-starter/database"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// MigrateCmd represents the migrate command
var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the database connection
		db := database.GetDb()

		// Ensure the migrations table exists
		createMigrationsTableIfNotExist(db)

		// Path to the migrations folder
		migrationsDir := "database/migrations"

		// Get all migration files from the migrations folder
		files, err := filepath.Glob(filepath.Join(migrationsDir, "*.go"))
		if err != nil {
			log.Fatalf("Error reading migration files: %v", err)
		}

		// Get the applied migrations
		var appliedMigrations []models.Migration
		if err := db.Find(&appliedMigrations).Error; err != nil {
			log.Fatalf("Error checking applied migrations: %v", err)
		}

		// Get the names of the applied migrations
		appliedMigrationNames := make(map[string]bool)
		for _, m := range appliedMigrations {
			appliedMigrationNames[m.Name] = true
		}

		// Loop through all migration files and run them
		for _, file := range files {
			// Extract migration name from the file
			fileName := filepath.Base(file)
			migrationName := strings.TrimSuffix(fileName, ".go")

			// If the migration has already been applied, skip it
			if appliedMigrationNames[migrationName] {
				fmt.Printf("Migration %s already applied, skipping...\n", migrationName)
				continue
			}

			// Dynamically execute the Up function of the migration
			runMigration(db, migrationName)

			// After running the migration, record it as applied
			if err := db.Create(&models.Migration{Name: migrationName}).Error; err != nil {
				log.Fatalf("Failed to record migration %s: %v", migrationName, err)
			}
		}

		fmt.Println("Migrations ran successfully!")
	},
}

// Helper function to check and create the migrations table if it doesn't exist
func createMigrationsTableIfNotExist(db *gorm.DB) {
	// Check if the migrations table exists
	if !db.Migrator().HasTable(&models.Migration{}) {
		// If it doesn't exist, create the migrations table
		if err := db.AutoMigrate(&models.Migration{}); err != nil {
			// log.Fatalf("Failed to create migrations table: %v", err)
		}
	}
}

// Helper function to run migration
func runMigration(db *gorm.DB, migrationName string) {
	fmt.Printf("Running migration: %s\n", migrationName)

	// Dynamically load the migration file
	migrationPackage := fmt.Sprintf("github.com/devndam/go-starter/database/migrations.%s", migrationName)

	// Reflectively find the Up function from the migration file
	upFunction := reflect.ValueOf(migrationPackage).MethodByName("Up")
	if !upFunction.IsValid() {
		log.Fatalf("Migration %s does not have an Up function", migrationName)
	}

	// Apply the migration (call the Up function)
	err := upFunction.Call([]reflect.Value{reflect.ValueOf(db)})
	if err != nil {
		log.Fatalf("Migration %s failed: %v", migrationName, err)
	}
}

package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// Global list of migratable models
var MigratableModels []interface{}

// CreateModelCmd represents the create:model command
var CreateModelCmd = &cobra.Command{
	Use:   "create:model [model_name]",
	Short: "Create a new model file with migration",
	Args:  cobra.ExactArgs(1), // Ensure exactly 1 argument is passed
	Run: func(cmd *cobra.Command, args []string) {
		modelName := args[0]
		fmt.Printf("Creating model and migration for: %s\n", modelName)

		// Generate file paths
		modelFilePath := "app/models/" + modelName + ".go"
		migrationFilePath := fmt.Sprintf("database/migrations/%s_create_%s_table.go", time.Now().Format("20060102150405"), modelName)

		// Create the model file
		createModelFile(modelFilePath, modelName)

		// Create the migration file
		createMigrationFile(migrationFilePath, modelName)

		fmt.Println("Model and migration created successfully!")
	},
}

// Helper function to create model file
func createModelFile(filePath, modelName string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Could not create model file: %v", err)
	}
	defer file.Close()

	content := `package models

import "time"

type ` + modelName + ` struct {
	ID        uint      ` + "`json:\"id\" gorm:\"primaryKey\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\"`" + `
	// Add fields for your model here
}
`
	file.WriteString(content)
}

// Helper function to create migration file
func createMigrationFile(filePath, modelName string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Could not create migration file: %v", err)
	}
	defer file.Close()

	content := `package migrations

import (
	"github.com/devndam/go-starter/app/models"
	"gorm.io/gorm"
)

func Up(db *gorm.DB) error {
	// Create the table for the ` + modelName + ` model
	return db.AutoMigrate(&models.` + modelName + `{})
}

func Down(db *gorm.DB) error {
	// Drop the table for the ` + modelName + ` model
	return db.Migrator().DropTable(&models.` + modelName + `{})
}
`
	file.WriteString(content)
}

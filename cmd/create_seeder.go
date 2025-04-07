package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var CreateSeederCmd = &cobra.Command{
	Use:   "create:seeder [SeederName]",
	Short: "Create a new seeder and register it",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		seederName := args[0]
		createSeederFile(seederName)
		registerSeederInRegistry("database/seeder_registry.go", seederName)
		fmt.Printf("Seeder %s created and registered successfully.\n", seederName)
	},
}

func createSeederFile(seederName string) {
	fileName := fmt.Sprintf("database/seeders/%s.go", strings.ToLower(seederName))
	if _, err := os.Stat(fileName); err == nil {
		fmt.Println("Seeder already exists.")
		return
	}

	content := fmt.Sprintf(`package seeders

		import (
			"github.com/devndam/go-starter/app/models"
			"gorm.io/gorm"
		)

		func %s(db *gorm.DB) error {
			// Example seeding logic
			return db.Create(&models.User{
				// Fill fields here
			}).Error
		}
`, seederName)

	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Could not create seeder file: %v", err)
	}
}

func registerSeederInRegistry(filePath, seederName string) {
	contentBytes, err := os.ReadFile(filePath)
	ref := "seeders." + seederName
	if err != nil {
		initial := fmt.Sprintf(`package database

import (
	"github.com/devndam/go-starter/database/seeders"
	"gorm.io/gorm"
)

type SeederFunc func(*gorm.DB) error

var SeederList = []SeederFunc{
	%s,
}
`, ref)
		os.WriteFile(filePath, []byte(initial), 0644)
		return
	}

	content := string(contentBytes)

	if strings.Contains(content, ref) {
		fmt.Println("Seeder already registered.")
		return
	}

	insertIndex := strings.LastIndex(content, "}")
	if insertIndex == -1 {
		log.Fatalf("Invalid SeederList format")
	}

	newContent := content[:insertIndex] + "\t" + ref + ",\n" + content[insertIndex:]
	os.WriteFile(filePath, []byte(newContent), 0644)
}

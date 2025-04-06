package main

import (
	"log"

	"os"

	"github.com/devndam/go-starter/app/config"
	"github.com/devndam/go-starter/cmd" // Import the cmd package to register commands
	"github.com/devndam/go-starter/database"
	"github.com/devndam/go-starter/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

// Root command
var rootCmd = &cobra.Command{
	Use:   "go-starter-cli",
	Short: "CLI for managing Go Starter project",
}

func init() {
	// Initialize environment variables
	config.LoadEnVariables()
	database.ConnectDb()
	// Register the subcommands for CLI
	// These commands are implemented in the cmd folder
	rootCmd.AddCommand(cmd.CreateModelCmd)
	rootCmd.AddCommand(cmd.MigrateCmd)
}

func main() {
	// Check if we are running a CLI command or starting the server
	if len(os.Args) > 1 {
		// If args are passed, execute the root command (CLI commands)
		if err := rootCmd.Execute(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	} else {
		// Start the server when no command is passed
		app := fiber.New()

		// Setup routes
		routes.ApiRoutes(app)

		// Start the server
		log.Fatal(app.Listen(":3000"))
	}
}

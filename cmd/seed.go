package cmd

import (
	"fmt"
	"log"

	"github.com/devndam/go-starter/database"
	"github.com/spf13/cobra"
)

var SeedCmd = &cobra.Command{
	Use:   "db:seed",
	Short: "Run database seeders",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.GetDb()

		for _, seeder := range database.SeederList {
			if err := seeder(db); err != nil {
				log.Fatalf("Seeder failed: %v", err)
			}
		}

		fmt.Println("All seeders ran successfully!")
	},
}

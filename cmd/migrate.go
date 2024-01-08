package cmd

import (
	"micgofiber/db"

	"github.com/spf13/cobra"
)

func InitMigrate() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate all schema",
	Long:  `All software has schemas.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Migrate()
	},
}

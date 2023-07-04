package cmd

import (
	"fmt"

	"github.com/remyz17/godoorpc/internal/utils"
	"github.com/spf13/cobra"
)

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var listDatabaseCmd = &cobra.Command{
	Use:   "list",
	Short: "List Odoo database",
	Long:  `This command show up the availible databases on Odoo server.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbs, err := odooService.ListDatabase()
		cobra.CheckErr(err)
		fmt.Println("List of databases:")
		for _, db := range *dbs {
			fmt.Println(db)
		}
	},
}

var databaseExistsCmd = &cobra.Command{
	Use:   "exists",
	Short: "Check if database exists",
	Long:  `This command check if the given database(s) exists in Odoo server.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbs, err := odooService.ListDatabase()
		cobra.CheckErr(err)
		for _, db := range args {
			_, found := utils.Find(*dbs, db)
			if !found {
				fmt.Println("Database", db, "does not exists")
			} else {
				fmt.Println("Database", db, "exists")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.AddCommand(listDatabaseCmd)
	databaseCmd.AddCommand(databaseExistsCmd)
}

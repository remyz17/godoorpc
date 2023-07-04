package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Odoo server manipulation",
	Long:  `Server command allow you to get various informations on the server`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Odoo version",
	Long: `By running this command, you will print server_serie attribute.
This attribute comes from common.version() rpc method`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := odooService.Version()
		cobra.CheckErr(err)
		fmt.Printf("Version: %s\n", result.ServerVersion)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(versionCmd)
}

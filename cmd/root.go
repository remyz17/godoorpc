package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "godoorpc",
	Short: "Odoo RPC CLI tool",
	Long: `GodooRPC allow you to communicate with Odoo ERP through it's external API.
	
Both XML and JSON RPC protocols are supported.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := viper.GetString("url")
		db := viper.GetString("db")
		username := viper.GetString("username")
		password := viper.GetString("password")

		fmt.Printf("URL: %s, DB: %s, Username: %s, Password: %s\n", url, db, username, password)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.godoorpc.yaml)")
	rootCmd.PersistentFlags().String("url", "", "URL du serveur Odoo")
	rootCmd.PersistentFlags().String("db", "", "Nom de la base de données Odoo")
	rootCmd.PersistentFlags().String("username", "", "Nom d'utilisateur Odoo")
	rootCmd.PersistentFlags().String("password", "", "Mot de passe Odoo")

	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("db", rootCmd.PersistentFlags().Lookup("db"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".godoorpc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".godoorpc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

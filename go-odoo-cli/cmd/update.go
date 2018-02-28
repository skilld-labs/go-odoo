package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [args]",
	Short: "update the given odoo model",
	Long:  "This command update the needed api and type files for the given odoo model",
	Example: `
	./go-odoo-cli update account.analytic.account
	./go-odoo-cli update account.analytic.account account.invoice
	./go-odoo-cli update all`,
	Run: upsert,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().StringP("uri", "", "http://localhost:8069", "the odoo instance URI")
	updateCmd.PersistentFlags().StringP("database", "d", "database", "the name of the postgres database linked to odoo")
	updateCmd.PersistentFlags().StringP("username", "u", "admin", "the odoo instance administrator")
	updateCmd.PersistentFlags().StringP("password", "p", "password", "the odoo instance administrator password")
}

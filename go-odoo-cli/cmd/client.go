package cmd

import (
	"../../api"
	"github.com/spf13/cobra"
)

func getClient(cmd *cobra.Command) (*api.Client, error) {
	uri, err := cmd.PersistentFlags().GetString("uri")
	if err != nil {
		return nil, err
	}
	db, err := cmd.PersistentFlags().GetString("database")
	if err != nil {
		return nil, err
	}
	admin, err := cmd.PersistentFlags().GetString("username")
	if err != nil {
		return nil, err
	}
	password, err := cmd.PersistentFlags().GetString("password")
	if err != nil {
		return nil, err
	}
	config := &api.Config{
		DbName:   db,
		Admin:    admin,
		Password: password,
		URI:      uri,
	}
	return api.NewClient(config)
}

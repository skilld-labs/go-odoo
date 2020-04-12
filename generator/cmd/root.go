package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	odoo "github.com/ahuret/go-odoo"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "./go-odoo -u admin -p admin -d odoo --url http://localhost:8069 --models crm.lead",
		Short: "Generates your odoo models for go-odoo golang library.",
		Long: ` 
Generates your odoo models for go-odoo golang library.
You can provide models name as arguments to specify what models to generate. By default, it generates all models of your odoo instance
		`,
		Run: func(cmd *cobra.Command, args []string) {
			var mm []string
			if models != "" {
				mm = strings.Split(models, ",")
			}
			if err := g.handleModels(mm); err != nil {
				handleError(err)
			}
		},
	}
	database   string
	admin      string
	password   string
	url        string
	noFmt      bool
	destFolder string
	models     string
	c          *odoo.Client
	t          *template.Template
	g          *generator
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleError(err)
	}
}

func init() {
	cobra.OnInitialize(initOdoo, initTemplate, initGenerator)

	rootCmd.PersistentFlags().StringVarP(&database, "database", "d", "odoo", "the instance database")
	rootCmd.PersistentFlags().StringVarP(&admin, "admin", "u", "admin", "the admin username")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "admin", "the admin password")
	rootCmd.PersistentFlags().StringVar(&url, "url", "http://localhost:8069", "the url of your odoo instance")
	rootCmd.PersistentFlags().StringVarP(&destFolder, "dest", "o", "", "the destination of generated models")
	rootCmd.PersistentFlags().StringVarP(&models, "models", "m", "", "the models you want to generate, separated by commas, empty means generate all")
	rootCmd.PersistentFlags().BoolVar(&noFmt, "no-fmt", false, "specify if you want to disable auto format of generated models")
}

func initGenerator() {
	g = NewGenerator(GeneratorConfiguration{Odoo: c, ModelTemplate: t, FormatModels: !noFmt, DestFolder: destFolder})
}

func initOdoo() {
	var err error
	if c, err = odoo.NewClient(&odoo.ClientConfig{
		Admin:    admin,
		Password: password,
		Database: database,
		URL:      url,
	}); err != nil {
		handleError(err)
	}
}

func initTemplate() {
	var err error
	if t, err = template.New("model.tmpl").ParseFiles("./generator/cmd/tmpl/model.tmpl"); err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

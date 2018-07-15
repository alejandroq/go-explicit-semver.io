package cmd

import (
	"github.com/alejandroq/go-explicit-semver.io/src/log"
	_templates "github.com/alejandroq/go-explicit-semver.io/src/templates"
	"github.com/spf13/cobra"
)

var input *string
var output *string

var templateCmd = &cobra.Command{
	Use:   "templates",
	Short: "Execute actions concerning Templates",
	Long: `This subcommand will allow you to remove, add and list Templates. 
	An input Template has an output that is becomes an auto generated assets 
	that can distribute your versioning to documents, etc automatically.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := _templates.Templates("list", *input, *output)
		if err != nil {
			log.Log("error", map[string]interface{}{
				"err": err.Error(),
			})
		}
	},
}

var templateAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a Template",
	Long:  `This subcommand will allow you to add a Template.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := _templates.Templates("add", *input, *output)
		if err != nil {
			log.Log("error", map[string]interface{}{
				"err": err.Error(),
			})
		}
	},
}

var templateRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a Template",
	Long:  `This subcommand will allow you to remove a Template.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := _templates.Templates("rm", *input, *output)
		if err != nil {
			log.Log("error", map[string]interface{}{
				"err": err.Error(),
			})
		}
	},
}

var templateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Templates",
	Long:  `This subcommand will allow you to list your Templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := _templates.Templates("list", *input, *output)
		if err != nil {
			log.Log("error", map[string]interface{}{
				"err": err.Error(),
			})
		}
	},
}

func init() {
	RootCmd.AddCommand(templateCmd)
	templateCmd.AddCommand(templateAddCmd)
	templateCmd.AddCommand(templateRmCmd)
	templateCmd.AddCommand(templateListCmd)

	input = templateCmd.Flags().String("input", "", "Template document")
	output = templateCmd.Flags().String("output", "", "Output to generate based on the input Template")
}

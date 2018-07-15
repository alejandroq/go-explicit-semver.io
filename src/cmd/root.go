package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-explicit-semver",
	Short: "Explicitly manage the semver of various files or directories",
	Long: `
Go-Explicit-Semver is an application that can manage the semantic versioning of an 
application via a ledger explicitly mutated by a user.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking for diffs...")
		fmt.Println("Incrementing Patches for...")
		fmt.Println("Complete :)")

		// use go routines per
		// Patch releases (diff)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-example.yaml)")
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Run verbosely")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("verbose", "v", false, "Run verbosely")
}

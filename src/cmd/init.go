package cmd

import (
	_init "github.com/alejandroq/go-explicit-semver.io/src/init"
	"github.com/alejandroq/go-explicit-semver.io/src/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// initCmd initializes v
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize v",
	Long: `This subcommand will generate the necessary files to accurately 
	maintain the semvers of your various files`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Elog.WithFields(logrus.Fields{
			"args": args,
		}).Info("arguments")

		err := _init.Init(args)
		if err != nil {
			log.Elog.WithFields(logrus.Fields{
				"err": err.Error(),
			}).Error("error")
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

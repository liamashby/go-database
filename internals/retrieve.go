package cmd

import (
	"github.com/spf13/cobra"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "retrieve data",
	Long:  `retrieve all of the data that was written to disk`,
	Run: func(cmd *cobra.Command, args []string) {
		testmodule.retrieve()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

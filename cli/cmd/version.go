package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of database cli",
	Long:  `All software has versions. This is database cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("the version number is 1.0.0")
	},
}

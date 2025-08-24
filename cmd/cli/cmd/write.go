package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func runWrite(cmd *cobra.Command, args []string) {
	fmt.Println("creating directory if it doesn't exist...")

	err := os.Mkdir("./data", 0750)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("creating a new file")
	file, err := os.Create("./data/text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString("test")
}

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "write data",
	Long:  `write data to disk using a cli`,
	Run:   runWrite,
}

func init() {
	rootCmd.AddCommand(writeCmd)
}

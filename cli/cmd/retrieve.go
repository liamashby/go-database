package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "retrieve data",
	Long:  `retrieve all of the data that was written to disk`,
	Run: func(cmd *cobra.Command, args []string) {
		read_index()
	},
}

type Mapping struct {
	ID   string `json:"id"`
	File string `json:"file"`
}

func read_index() {
	jsonFile, err := os.Open("./data/index.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var mappings []Mapping
	json.Unmarshal(byteValue, &mappings)
}

func init() {
	rootCmd.AddCommand(retrieveCmd)
}

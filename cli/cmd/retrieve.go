package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

type Mapping struct {
	ID   string `json:"id"`
	File string `json:"file"`
}

func read_and_filter_index(args []string) []string {
	jsonFile, err := os.Open("./data/index.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var mappings []Mapping
	json.Unmarshal(byteValue, &mappings)

	var files_to_read []string
	for i := 0; i < len(mappings); i++ {
		if slices.Contains(args, mappings[i].ID) {
			files_to_read = append(files_to_read, mappings[i].File)
		}
	}
	return files_to_read
}

func read_data_on_disk(files []string) {
}

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "retrieve data",
	Long:  `retrieve all of the data that was written to disk`,
	Run: func(cmd *cobra.Command, args []string) {
		read_and_filter_index(args)
	},
}

func init() {
	rootCmd.AddCommand(retrieveCmd)
}

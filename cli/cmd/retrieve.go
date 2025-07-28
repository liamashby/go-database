package cmd

import (
	"fmt"
	"os"
)

func retrieve(file_name string) error {
	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	return nil
}

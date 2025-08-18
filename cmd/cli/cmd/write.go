package cmd

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("creating directory if it doesn't exist")

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

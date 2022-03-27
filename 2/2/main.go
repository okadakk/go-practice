package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{"1", "Java"})
	writer.Write([]string{"2", "Golang"})
	writer.Write([]string{"3", "Ruby"})
	writer.Flush()
}

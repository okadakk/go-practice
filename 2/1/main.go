package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(file, "show number %d\n", 123)
	fmt.Fprintf(file, "show string %s\n", "hello world!")
	fmt.Fprintf(file, "show float %f\n", 1.23)

	file.Close()
}

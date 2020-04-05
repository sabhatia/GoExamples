package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Process the name of the file from the arg line
	fmt.Println("Passed in", len(os.Args), "arguments")
	for i, arg := range os.Args {
		fmt.Println("Arg[", i, "] = ", arg)
	}

	// Use the first argument as the filename
	fmt.Println("Opening file:", os.Args[1])
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
	file.Close()
}

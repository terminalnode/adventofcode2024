package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to read file '%s' with error: %v\n", path, err)
		os.Exit(1)
	}

	buff := bytes.NewBuffer(data)
	input, err := io.ReadAll(buff)
	if err != nil {
		fmt.Printf("Failed to parse data from file '%s' with error: %v\n", path, err)
	}

	return string(input)
}

package main

import (
	"io"
	"os"
)

func main() {
	filename := os.Args[1]

	f, err := os.Open(filename)

	if err != nil {
		os.Exit(2)
	}

	io.Copy(os.Stdout, f)
}
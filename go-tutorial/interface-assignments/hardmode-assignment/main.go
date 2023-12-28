package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args)
	filename := os.Args[1]
	
	f, err := os.Open(filename)

	if err != nil {
		os.Exit(2)
	}

	io.Copy(os.Stdout, f)
}
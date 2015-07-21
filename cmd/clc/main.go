package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	output := run(args)
	fmt.Printf("%s\n", output)
}

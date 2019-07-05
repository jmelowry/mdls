package main

import (
	"fmt"
	"os"
)

func main() {
	allArgs := os.Args[1:]
	fmt.Println(allArgs)
}

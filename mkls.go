package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.Create("list.md")
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))

	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	if len(os.Args) > 1 {
		fmt.Println("2 or more args")
	}

	//d := os.Args

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("list created?")
}

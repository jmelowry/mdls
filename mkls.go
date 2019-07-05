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

	allArgs := os.Args

	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	//d := allArgs

	//fmt.Fprintln(d)

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("list created?")
}

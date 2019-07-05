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

	//d := []string{allArgs}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("list created?")
}

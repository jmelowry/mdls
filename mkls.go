package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("list.md")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	d := []string{"This is my simple(?) list."}

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

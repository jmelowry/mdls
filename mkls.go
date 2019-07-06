package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("mkls", "1.0.0")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"foo": fooCommandFactory,
		"bar": barCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)

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

func initialize() {
	fmt.Println("Initialized")
}

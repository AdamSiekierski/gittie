package main

import (
	"fmt"
	"os"
	"log"
	"github.com/urfave/cli/v2"
)

func application(c *cli.Context) error {
	fmt.Println("hello world!")
	return nil
}

func main() {
	app := &cli.App{
		Name: "Gittie",
		Usage: "Quickly create .gitignore in your project using one simple command ",
		Action: application,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatalln(err)
	}
}

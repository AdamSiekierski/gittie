package main

import (
	"fmt"
	"os"
	"log"
	"github.com/urfave/cli/v2"
	"github.com/AdamSiekierski/gittie/utils"
)

func application(c *cli.Context) error {
	if c.Bool("list") {
		_, err := utils.GetTemplates()

		if err != nil {
			log.Fatalln(err)
		}
	}

	if c.String("create") != "" {
		fmt.Printf("hmmm %s", c.String("create"))
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "Gittie",
		Usage: "Quickly create .gitignore in your project using one simple command ",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "list",
				Usage:   "Show all available templates",
				Aliases: []string{"l"},
				Value:   false,
			},
			&cli.StringFlag{
				Name:    "template",
				Usage:   "Generate a .gitignore from selected template",
				Aliases: []string{"t"},
				Value:   "",
			},
		},
		HideHelpCommand: true,
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool("list") {
			fmt.Println("list")
		}

		if c.String("template") != "" {
			fmt.Printf("template %s", c.String("create"))
		}

		return nil
	}

	app.Run(os.Args)
}

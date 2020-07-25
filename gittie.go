package main

import (
	"fmt"
	"github.com/AdamSiekierski/gittie/utils"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

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
			templates, err := utils.GetTemplates()

			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(templates)
		}

		if c.String("template") != "" {
			fmt.Printf("template %s", c.String("create"))
		}

		return nil
	}

	app.Run(os.Args)
}

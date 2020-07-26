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
			templates, err := utils.AllTemplates()

			if err != nil {
				log.Fatalln(err)
				return err
			}

			fmt.Printf("ALL TEMPLATES: \n\n")

			for i, v := range templates {
				if (i % 2) != 0 {
					fmt.Printf("|%-21s|\n", v)
				} else {
					fmt.Printf("|%-21s|", v)
				}
			}
		}

		if c.String("template") != "" {
			template, err := utils.GetTemplate(c.String("template"))

			if err != nil {
				log.Fatalln(err)
				return err
			}

			err = utils.WriteGitignore(template)

			if err != nil {
				log.Fatalln(err)
				return err
			}

			fmt.Printf("Successfully created .gitignore from template: %s \n", template["name"])
		}

		return nil
	}

	app.Run(os.Args)
}

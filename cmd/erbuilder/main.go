package cmd

import (
	"log"
	"os"

	"github.com/Angelos-Giannis/erbuilder/internal/app/service"
	"github.com/Angelos-Giannis/erbuilder/internal/config"
	"github.com/Angelos-Giannis/erbuilder/internal/domain"
	"github.com/Angelos-Giannis/erbuilder/internal/pkg/util"
	"github.com/urfave/cli/v2"
)

// Main actual main function.
func Main() {
	cfg := config.New()

	var app = cli.NewApp()
	info(app, cfg)

	options := domain.NewOptions(cfg)

	app.Commands = []*cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generate the .er file based on the provided structures.",
			Flags: []cli.Flag{
				options.GetCommonFields(),
				options.GetDirectoryFlag(),
				options.GetFileList(),
				options.GetIDField(),
				options.GetOutputFilename(),
				options.GetOutputPath(),
				options.GetTag(),
				options.GetTitle(),
				options.GetColumnNameCase(),
				options.GetTableNameCase(),
				options.GetTableNamePlural(),
			},
			Action: func(c *cli.Context) error {
				err := options.Validate()
				if err != nil {
					panic(err)
				}

				util := util.New()

				srv := service.New(options, util)
				return srv.Generate()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// info sets up the information of the tool.
func info(app *cli.App, cfg config.Config) {
	var appAuthors []*cli.Author
	for _, author := range cfg.Application.Authors {
		newAuthor := cli.Author{
			Name:  author.Name,
			Email: author.Email,
		}
		appAuthors = append(appAuthors, &newAuthor)
	}

	app.Authors = appAuthors
	app.Name = cfg.Application.Name
	app.Usage = cfg.Application.Usage
	app.Version = cfg.Application.Version
}
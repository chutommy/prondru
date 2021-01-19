package controller

import (
	"time"

	"github.com/urfave/cli/v2"
)

// author stores author's contacts.
var author = &cli.Author{
	Name:  "Tommy Chu",
	Email: "tommychu2256@gmail.com",
}

// flags defines CLI's flags.
var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:        "id",
		Aliases:     []string{"i"},
		Usage:       "make ID field required",
		Value:       false,
		DefaultText: "false",
	},
	&cli.BoolFlag{
		Name:        "author",
		Aliases:     []string{"a"},
		Usage:       "make Author field required",
		Value:       false,
		DefaultText: "false",
	},
	&cli.BoolFlag{
		Name:        "title",
		Aliases:     []string{"i"},
		Usage:       "make Title field required",
		Value:       false,
		DefaultText: "false",
	},
	&cli.IntFlag{
		Name:        "results",
		Aliases:     []string{"r"},
		Usage:       "max number of results",
		Value:       60,
		DefaultText: "60",
	},
}

// NewApp returns the prondru CLI application.
func NewApp() *cli.App {
	return &cli.App{
		Name:        "prOndru",
		Usage:       "CLI for Office of Scientific and Technical Information API",
		Version:     "18.0.0",
		Description: "Terminal gateway to the fantastic world of physical discoveries.",
		Flags:       flags,
		Action:      action,
		Compiled:    time.Time{},
		Authors:     []*cli.Author{author},
	}
}

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

var (
	idFlag      = "id"
	authorFlag  = "author"
	titleFlag   = "flag"
	resultsFlag = "results"
)

// flags defines CLI's flags.
var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:        idFlag,
		Aliases:     []string{"i"},
		Usage:       "make ID field required",
		Value:       false,
		DefaultText: "false",
	},
	&cli.BoolFlag{
		Name:        authorFlag,
		Aliases:     []string{"a"},
		Usage:       "make Author field required",
		Value:       false,
		DefaultText: "false",
	},
	&cli.BoolFlag{
		Name:        titleFlag,
		Aliases:     []string{"t"},
		Usage:       "make Title field required",
		Value:       false,
		DefaultText: "false",
	},
	&cli.IntFlag{
		Name:        resultsFlag,
		Aliases:     []string{"r"},
		Usage:       "`MAX` number of results",
		Value:       100,
		DefaultText: "100",
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

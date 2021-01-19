package controller

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"prondru/api"
	"prondru/data"
	"prondru/handler"
	"prondru/show"
	"time"
)

// action defines the process of the CLI.
func action(c *cli.Context) error {
	// retrieve flags
	byID := c.Bool(idFlag)
	byAuthor := c.Bool(authorFlag)
	byTitle := c.Bool(titleFlag)
	rows := c.Int(resultsFlag)

	// get configuration and construct handler
	cfg := handler.NewConfig(byID, byAuthor, byTitle, rows)
	h := handler.NewHandler(api.URL, cfg)

	// do cycle
	for {
		if err := cycle(h); err != nil {
			return err
		}
	}
}

// cycle: prompt user -> fetch data -> let user select -> printout the selection.
func cycle(h *handler.Handler) error {
	// prompt
	r, err := h.Prompt()
	if err != nil {
		return err
	}

	// fetch
	rr, err := data.Fetch(h.URL(), r)
	if err != nil {
		return err
	}

	// select
	resp, err := h.Select(rr)
	if err != nil {
		return err
	}

	// print
	show.PrintResponse(resp)

	// wait
	fmt.Printf("Press enter to continue...\n")
	time.Sleep(250 * time.Millisecond)
	fmt.Scanln()

	return nil
}

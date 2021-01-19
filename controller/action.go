package controller

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"prondru/api"
	"prondru/data"
	"prondru/handler"
	"prondru/show"
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
	fmt.Println("++++++++++++++++")

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

	if len(rr) == 0 {
		fmt.Println()
		fmt.Println("It looks like there aren't many great articles for your search...")
		fmt.Println("Try something else :)")

		return nil
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
	bufio.NewReader(os.Stdin).ReadByte() // wait enter
	fmt.Println()

	return nil
}

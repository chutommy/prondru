package controller

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"prondru/api"
	"prondru/data"
	"prondru/handler"
	"prondru/show"

	"github.com/dimiro1/banner"
	_ "github.com/dimiro1/banner/autoload" // load banner
	"github.com/mattn/go-colorable"
	"github.com/urfave/cli/v2"
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

	// print banner
	banner.Init(colorable.NewColorableStdout(), true, false, bytes.NewBufferString(show.Banner()))
	fmt.Println()

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
		return fmt.Errorf("could not retrieve prompt input: %w", err)
	}

	// fetch
	rr, err := data.Fetch(h.URL(), r)
	if err != nil {
		return fmt.Errorf("could not fetch url: %w", err)
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
		return fmt.Errorf("could not get the selection: %w", err)
	}

	// print
	show.PrintResponse(resp)

	// wait
	fmt.Printf("Press enter to continue...\n")

	_, _ = bufio.NewReader(os.Stdin).ReadByte()

	fmt.Println()

	return nil
}

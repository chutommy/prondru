package handler

import (
	"fmt"
	"prondru/data"

	"github.com/manifoldco/promptui"
)

// selectRecord asks user to choose one of the articles.
func selectRecord(records data.Responses) (*data.Response, error) {
	selection := promptui.Select{
		Label: "Select article",
		Items: records.ToStrings(),
		Size:  10,
	}

	// run the selection
	i, _, err := selection.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to rertieve a selection: %w", err)
	}

	return records[i], nil
}

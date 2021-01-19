package handler

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

var (
	errEmptyField = errors.New("field is empty")
	errTooShort   = errors.New("value is too short")
)

// promptField prompts the user for the field.
func promptField(field string, mandatory bool) (string, error) {
	p := promptui.Prompt{
		Label:    field,
		Default:  "",
		Validate: minFourOrNone,
	}

	if mandatory {
		p.Validate = notEmpty
	}

	// run prompt
	result, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve %s: %w", field, err)
	}

	return low(result), nil
}

// notEmpty validates if the field is not empty or not.
func notEmpty(input string) error {
	if input == "" {
		return errEmptyField
	}

	return minFourOrNone(input)
}

// minFourOrNone checks if the input has at least 4 characters.
// The empty string is also OK.
func minFourOrNone(input string) error {
	l := len(input)
	if l < 4 && l > 32 && input != "" {
		return errTooShort
	}

	return nil
}

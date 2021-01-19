package data

import (
	"encoding/json"
	"fmt"
	"io"
)

// Response represents a search result.
type Response struct {
	Id                 string   `json:"osti_id"`
	Title              string   `json:"title"`
	Authors            []string `json:"authors"`
	Subjects           []string `json:"subjects"`
	Description        string   `json:"description"`
	Doi                string   `json:"doi"`
	Publisher          string   `json:"publisher"`
	CountryPublication string   `json:"country_publication"`
	PublicationDate    string   `json:"publication_date"`
}

// Responses represents a list of Response.
type Responses []*Response

// toResponses decodes the io.Reader into Responses.
func toResponses(inp io.Reader) (Responses, error) {
	var responses Responses

	// decode
	err := json.NewDecoder(inp).Decode(&responses)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return responses, nil
}

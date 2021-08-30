package handler

import (
	"github.com/chutommy/prondru/data"
)

// Handler handles user communication.
type Handler struct {
	apiURL string
	cfg    *Config
}

// NewHandler is a constructor for the Handler struct.
func NewHandler(u string, cfg *Config) *Handler {
	return &Handler{
		apiURL: u,
		cfg:    cfg,
	}
}

// Prompt prompts the user based on the Handler's Config.
func (h *Handler) Prompt() (*data.Request, error) {
	var id, author, title, query string

	var err error

	// query optional fields

	// if id, err = promptField("ID", h.cfg.byID); err != nil {
	// 	return nil, err
	// }

	// if author, err = promptField("Author", h.cfg.byAuthor); err != nil {
	// 	return nil, err
	// }

	// if title, err = promptField("Title", h.cfg.byTitle); err != nil {
	// 	return nil, err
	// }

	// decide if query is mandatory
	// b := true
	// if h.cfg.byID || h.cfg.byAuthor || h.cfg.byTitle {
	// 	b = false
	// }

	if query, err = promptField("Keyword", true); err != nil {
		return nil, err
	}

	// construct a new request
	return data.NewRequest(id, query, author, title, h.cfg.rows, h.cfg.page), nil
}

// Select prompts user to select one of the Responses.
func (h *Handler) Select(rr data.Responses) (*data.Response, error) {
	return selectRecord(rr)
}

// URL returns apiURL.
func (h *Handler) URL() string {
	return h.apiURL
}

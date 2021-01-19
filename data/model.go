package data

// Request represents a search query configuration.
type Request struct {
	id     string
	query  string
	author string
	title  string
	rows   int
	page   int
}

// Response represents a search result.
type Response struct {
	id                 string
	title              string
	authors            []string
	subjects           []string
	description        string
	doi                string
	publisher          string
	countryPublication string
	publicationDate    string
}

// Results represents a list of Response.
type Results []*Response

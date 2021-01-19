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

// NewRequest constructs a new Request.
func NewRequest(id, q, author, title string, rows, page int) *Request {
	return &Request{
		id:     id,
		query:  q,
		author: author,
		title:  title,
		rows:   rows,
		page:   page,
	}
}

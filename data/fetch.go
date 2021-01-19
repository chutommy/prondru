package data

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Fetch makes a request based on the given Request and
// returns Responses.
func Fetch(u string, r *Request) (Responses, error) {
	// get a url
	newURL := getURL(u, r)

	// get response
	resp, err := http.Get(newURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make a reuqest: %w", err)
	}

	// defer body closure
	defer func() {
		_ = resp.Body.Close()
	}()

	// decode
	responses, err := toResponses(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not unparse response's body content: %w", err)
	}

	return responses, nil
}

// getURL merges Request's data into the URL's query string.
func getURL(u string, r *Request) string {
	newURL, _ := url.Parse(u)
	q := url.Values{}

	// add id
	if r.id != "" {
		//goland:noinspection SpellCheckingInspection
		q.Add("osti_id", r.id)
	}

	// add query
	if r.query != "" {
		q.Add("q", r.query)
	}

	// add author
	if r.author != "" {
		q.Add("author", r.author)
	}

	// add title
	if r.title != "" {
		q.Add("title", r.title)
	}

	// add pagination data
	q.Add("rows", strconv.Itoa(r.rows))
	q.Add("page", strconv.Itoa(r.page))

	// encode queries
	newURL.RawQuery = q.Encode()

	return newURL.String()
}

package show

import (
	"fmt"
	"prondru/data"
	"strings"

	"github.com/fatih/color"
)

// PrintResponse prints formatted article.
//goland:noinspection ALL
func PrintResponse(r *data.Response) {
	b := color.New(color.Bold)
	u := color.New(color.Underline)
	i := color.New(color.Italic)
	f := color.New(color.Faint)

	fmt.Printf("\n")
	fmt.Printf("================\n")

	// {title} [OSTI: {id}]
	b.Printf("%s", r.Title)
	fmt.Printf(" ")
	u.Printf("[OSTI: %s]", r.ID)
	fmt.Printf("\n")

	// Authors: {authors}
	// Publisher: {publisher} - {country}
	b.Printf("Authors: ")
	i.Printf("%s\n", join(r.Authors))
	b.Printf("Publisher: ")
	fmt.Printf("%s - %s\n", r.Publisher, r.CountryPublication)

	// Subjects: {subjects}
	b.Printf("Subjects: ")
	fmt.Printf("%s\n", low(join(r.Subjects)))

	// Date: {publication data} ({url})
	b.Printf("Date: ")
	f.Printf("%s (%s)\n", r.PublicationDate, r.Doi)

	fmt.Printf("----------------\n")

	// Research
	// {description}
	u.Printf("Research\n")
	fmt.Printf("%s\n", r.Description)
	fmt.Printf("\n")
}

// join joins slice of strings.
func join(ss []string) string {
	if len(ss) >= 7 {
		ss = ss[:7]
	}

	return strings.Join(ss, ", ")
}

// low returns the given string in lower-case.
func low(s string) string {
	return strings.ToLower(s)
}

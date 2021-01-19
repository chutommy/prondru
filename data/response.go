package data

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

// Results represents a list of Response.
type Results []*Response

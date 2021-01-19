package handler

// Config represents a configuration of the current runtime.
type Config struct {
	byID     bool
	byAuthor bool
	byTitle  bool
	rows     int
	page     int
}

// NewConfig is a constructor of the Config.
func NewConfig(byID, byAuthor, byTitle bool, rows int) *Config {
	return &Config{
		byID:     byID,
		byAuthor: byAuthor,
		byTitle:  byTitle,
		rows:     rows,
		page:     1,
	}
}

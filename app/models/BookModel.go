package models

// BookModel represents a book in the library
type BookModel struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Year          int     `json:"year"`
	Rating        float64 `json:"rating"`
	Author        string  `json:"author"`
	Publisher     string  `json:"publisher"`
	PublishedYear int     `json:"published_year"`
	ISBN          string  `json:"isbn"`
	Genre         string  `json:"genre"`
	Description   string  `json:"description"`
	CoverImage    string  `json:"cover_image"`

	// Add other fields as needed
}

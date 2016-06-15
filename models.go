package main

// Blog -- tumblr blog metadata
type Blog struct {
	Title string `json:"title"`
	Posts int    `json:"posts"`
	Name  string `json:"name"`
}

// Post -- a tumblr post
type Post struct {
	BlogName  string   `json:"blog_name"`
	ID        uint32   `json:"id"`
	PostURL   string   `json:"post_url"`
	Slug      string   `json:"slug"`
	Type      string   `json:"type"`
	State     string   `json:"state"`
	Format    string   `json:"format"`
	ReblogKey string   `json:"reblog_key"`
	Tags      []string `json:"tags"`
	ShortURL  string   `json:"short_url"`
	NoteCount uint32   `json:"note_count"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Summary   string   `json:"summary"`
	Photos    struct {
		Caption        string  `json:"caption"`
		AlternateSizes []Photo `json:"alt_sizes"`
		Original       Photo   `json:"original_size"`
	} `json:"photos"`
}

// Photo -- photo object in tumblr feed
type Photo struct {
	URL    string `json:"url"`
	Width  uint32 `json:"width"`
	Height uint32 `json:"height"`
}

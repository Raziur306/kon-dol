package model

type NewsResponse struct {
	Total int        `json:"total"`
	Items []NewsItem `json:"items"`
}

type NewsItem struct {
	Headline       string `json:"headline"`
	URL            string `json:"url"`
	Slug           string `json:"slug"`
	AuthorName     string `json:"author-name"`
	LastPublished  int64  `json:"last-published-at"`
	HeroImageS3Key string `json:"hero-image-s3-key"`
	TrackId        string `json:"id"`
	Tags           []Tag  `json:"tags"`
}

type Tag struct {
	Properties struct {
		Images          []interface{} `json:"images"`
		MetaTitle       string        `json:"meta-title"`
		MetaKeywords    string        `json:"meta-keywords"`
		MetaDescription string        `json:"meta-description"`
	} `json:"properties"`
	Slug         string      `json:"slug"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	TagType      string      `json:"tag-type"`
	EntityTypeID int         `json:"entity-type-id"`
	ExternalID   interface{} `json:"external-id"` // could be string or null, using interface{} for safety
	ID           int         `json:"id"`
}

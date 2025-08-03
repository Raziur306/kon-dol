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
}

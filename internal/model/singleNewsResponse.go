package model

type SingleNewsResponse struct {
	DomainSlug     *string     `json:"domainSlug"`
	CurrentHostURL string      `json:"currentHostUrl"`
	PrimaryHostURL string      `json:"primaryHostUrl"`
	HTTPStatusCode int         `json:"httpStatusCode"`
	PageType       string      `json:"pageType"`
	Data           ArticleData `json:"data"`
}

type ArticleData struct {
	Story Story `json:"story"`
}

type Story struct {
	ID               string        `json:"id"`
	LastUpdatedAt    int64         `json:"last-published-at"`
	Headline         string        `json:"headline"`
	Subheadline      string        `json:"subheadline"`
	Slug             string        `json:"slug"`
	URL              string        `json:"url"`
	AuthorName       string        `json:"author-name"`
	Sections         []Section     `json:"sections"`
	Authors          []Author      `json:"authors"`
	Summary          string        `json:"summary"`
	ContentType      string        `json:"content-type"`
	PublishedAt      int64         `json:"published-at"`
	HeroImageCaption *string       `json:"hero-image-caption"`
	Cards            []Card        `json:"cards"`
	Metadata         StoryMetadata `json:"metadata"`
}

type Section struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	SectionURL  string `json:"section-url"`
	DisplayName string `json:"display-name"`
}

type Author struct {
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar-url"`
}

type Card struct {
	StoryElements []StoryElement `json:"story-elements"`
}

type StoryElement struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type StoryMetadata struct {
	Excerpt string `json:"excerpt"`
}

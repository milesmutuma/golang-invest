package news

import "time"

// Article represents a news article.
type Article struct {
	PublishOn time.Time // PublishOn represents the publication date and time of the article.
	Headline  string    // Headline represents the title of the article.
}

// Fetcher is an interface for fetching news articles.
type Fetcher interface {
	Fetch(query string) ([]Article, error)
}

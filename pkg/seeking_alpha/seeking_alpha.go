package seeking_alpha

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"opg-analysis/internal/news"
)

const (
	urlPath      = "/news/v2/list-by-symbol"
	apiKeyHeader = "x-rapidapi-key"
	pageSize     = 5
)

// client is a struct that represents the Seeking Alpha API client
type client struct {
	apiKey  string
	baseUrl string
}

func (c *client) Fetch(ticker string) ([]news.Article, error) {
	// Build the URL for the API request
	url, err := c.buildURL(ticker)
	if err != nil {
		return nil, err
	}
	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add the API key header to the request
	req.Header.Add(apiKeyHeader, c.apiKey)

	// Send the HTTP request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Check if the response status code is successful
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("unsuccessful status code %d received", resp.StatusCode)
	}

	// Parse the response and extract the articles
	return c.parse(resp)
}

// parse decodes the JSON response and extracts the articles
func (c *client) parse(resp *http.Response) ([]news.Article, error) {
	res := &SeekingAlphaResponse{}

	err := json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, err
	}

	var articles []news.Article
	for _, item := range res.Data {
		art := news.Article{
			PublishOn: item.Attributes.PublishOn,
			Headline:  item.Attributes.Title,
		}
		articles = append(articles, art)
	}

	return articles, nil
}

// buidUrl contructs the URL for the API request
func (c *client) buildURL(ticker string) (string, error) {
	// Parse base url
	parsedUrl, err := url.Parse(c.baseUrl)
	if err != nil {
		return "", err
	}
	parsedUrl.Path += urlPath

	// set query parameters
	params := url.Values{}
	params.Add("size", fmt.Sprint(pageSize))
	params.Add("id", ticker)
	parsedUrl.RawQuery = params.Encode()

	return parsedUrl.String(), nil
}

func NewClient(baseURL, apiKey string) news.Fetcher {
	return &client{
		apiKey:  apiKey,
		baseUrl: baseURL,
	}
}

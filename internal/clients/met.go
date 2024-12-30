package clients

import "fmt"

type MetropolitanMuseumOfArtAPIClient struct {
	version    string
	base       string
	collection string
}

func NewMetropolianMuseumOfArtAPIClient() MetropolitanMuseumOfArtAPIClient {
	return MetropolitanMuseumOfArtAPIClient{
		base:       "https://collectionapi.metmuseum.org",
		version:    "v1",
		collection: "public/collection",
	}
}

func (c *MetropolitanMuseumOfArtAPIClient) getBaseUrl() string {
	return fmt.Sprintf("%s/%s/%s", c.base, c.collection, c.version)
}

func (c *MetropolitanMuseumOfArtAPIClient) GetObjectIDs() []string {
	url := fmt.Sprintf("%s/%s", c.getBaseUrl(), "objects")
	fmt.Println(url)
	return []string{}
}

func (c *MetropolitanMuseumOfArtAPIClient) GetObject(id string) any {
	url := fmt.Sprintf("%s/%s/%s", c.getBaseUrl(), "objects", id)
	fmt.Println(url)
	return nil
}

func (c *MetropolitanMuseumOfArtAPIClient) SearchForObject(query string) any {
	url := fmt.Sprintf("%s/%s?q=%s&hasImages=true", c.getBaseUrl(), "search", query)
	fmt.Println(url)
	return nil
}

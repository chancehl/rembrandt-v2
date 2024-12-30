package met

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

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

// Gets all ObjectIDs in the MET API collection
func (c *MetropolitanMuseumOfArtAPIClient) GetObjectIDs() (*GetObjectsResponse, error) {
	url, _ := url.JoinPath(c.base, []string{c.collection, c.version, "objects"}...)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not fetch object IDs from MET API: %v", err)
	}
	defer resp.Body.Close()

	var getObjectsResponse GetObjectsResponse
	if err := json.NewDecoder(resp.Body).Decode(&getObjectsResponse); err != nil {
		return nil, fmt.Errorf("could not deserialize MET API body: %v", err)
	}

	return &getObjectsResponse, nil
}

// Gets a MET object by ID
func (c *MetropolitanMuseumOfArtAPIClient) GetObjectByID(id string) (*GetObjectResponse, error) {
	url, _ := url.JoinPath(c.base, []string{c.collection, c.version, "objects", id}...)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not fetch object by ID from MET API: %v", err)
	}
	defer resp.Body.Close()

	var getObjectResponse GetObjectResponse
	if err := json.NewDecoder(resp.Body).Decode(&getObjectResponse); err != nil {
		return nil, fmt.Errorf("could not deserialize MET API body: %v", err)
	}

	return &getObjectResponse, nil
}

// Searches for objects matching query in MET API
func (c *MetropolitanMuseumOfArtAPIClient) SearchForObject(query string) (*GetObjectsResponse, error) {
	url, _ := url.JoinPath(c.base, []string{c.collection, c.version, "search"}...)
	urlWithQueryParams := fmt.Sprintf("%s?hasImages=true&q=%s", url, query)

	resp, err := http.Get(urlWithQueryParams)
	if err != nil {
		return nil, fmt.Errorf("could not query for objects in MET API: %v", err)
	}
	defer resp.Body.Close()

	var getObjectsResponse GetObjectsResponse
	if err := json.NewDecoder(resp.Body).Decode(&getObjectsResponse); err != nil {
		return nil, fmt.Errorf("could not deserialize MET API body: %v", err)
	}

	return &getObjectsResponse, nil
}

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/chancehl/rembrandt-v2/internal/cache"
)

type METApiClient struct {
	version    string
	base       string
	collection string
	cache      *cache.InMemoryCache
}

func NewMETApiClient(cache *cache.InMemoryCache) *METApiClient {
	return &METApiClient{
		base:       "https://collectionapi.metmuseum.org",
		version:    "v1",
		collection: "public/collection",
		cache:      cache,
	}
}

// Gets all ObjectIDs in the MET API collection
func (c *METApiClient) GetObjectIDs() (*GetObjectsResponse, error) {
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
func (c *METApiClient) GetObjectByID(id string) (*GetObjectResponse, error) {
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
func (c *METApiClient) SearchForObject(query string) (*GetObjectsResponse, error) {
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

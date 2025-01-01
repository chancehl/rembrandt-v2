package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/chancehl/rembrandt-v2/internal/cache"
)

const (
	ObjectIDsCacheKey = "objectIDs"
	ObjectIDsTTL      = time.Hour
)

type METAPIClient struct {
	version    string
	base       string
	collection string
	cache      *cache.InMemoryCache
}

func NewMETAPIClient(cache *cache.InMemoryCache) *METAPIClient {
	return &METAPIClient{
		base:       "https://collectionapi.metmuseum.org",
		version:    "v1",
		collection: "public/collection",
		cache:      cache,
	}
}

// Gets all ObjectIDs in the MET API collection
func (c *METAPIClient) GetObjectIDs() (*GetObjectsResponse, error) {
	if cachedObjectIDs, exists := c.cache.Get(ObjectIDsCacheKey); exists {
		if objectIDs, ok := cachedObjectIDs.([]int); ok {
			return &GetObjectsResponse{Total: len(objectIDs), ObjectIDs: objectIDs}, nil
		} else {
			return nil, fmt.Errorf("could not convert cached objectIDs to []int")
		}
	}

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

	// populate cache
	c.cache.Set(ObjectIDsCacheKey, getObjectsResponse.ObjectIDs, ObjectIDsTTL)

	return &getObjectsResponse, nil
}

// Gets a MET object by ID
func (c *METAPIClient) GetObjectByID(id int) (*GetObjectResponse, error) {
	url, _ := url.JoinPath(c.base, []string{c.collection, c.version, "objects", strconv.Itoa(id)}...)

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
func (c *METAPIClient) SearchForObject(query string) (*GetObjectsResponse, error) {
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

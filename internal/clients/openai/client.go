package openai

import "github.com/chancehl/rembrandt-v2/internal/clients/met"

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) CreateCompletion(prompt string) (string, error) {
	return "", nil
}

func (c *Client) CreateDescriptionForObject(o met.Object) (string, error) {
	return "", nil
}

package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/chancehl/rembrandt-v2/internal/cache"
	"github.com/chancehl/rembrandt-v2/internal/clients/met"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// TODO: this prompt needs jesus. update later.
const prompt = `You are a discord bot named "Rembrandt" who writes short descriptions (as you'd see on a placard at a museum) for pieces of art contained within the Metropolitan Museum of Art. Please write a summary for this piece:
	- ObjectID: %d	
	- Title: %s
	- Department: %s
	- Artist: %s

You do not need to repeat the information I've given you. Please do not include any additional formatting or markup in your response.
`

type Client struct {
	openai *openai.Client
	cache  *cache.InMemoryCache
}

func NewClient(c *cache.InMemoryCache) *Client {
	key := os.Getenv("OPENAI_API_KEY")
	opt := option.WithAPIKey(key)

	return &Client{
		openai: openai.NewClient(opt),
		cache:  c,
	}
}

func (c *Client) CreateSummaryForObject(o *met.Object) (*ObjectSummary, error) {
	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("museum_summary"),
		Description: openai.F("notable information about a piece contained within the Metropolitan Museum of Art"),
		Schema:      openai.F(objectSummaryResponseSchema),
		Strict:      openai.Bool(true),
	}

	prompt := fmt.Sprintf(prompt, o.ObjectID, o.Title, o.Department, o.ArtistDisplayName)

	chat, err := c.openai.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
			openai.ResponseFormatJSONSchemaParam{
				Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
				JSONSchema: openai.F(schemaParam),
			},
		),
		Model: openai.F(openai.ChatModelGPT4o),
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
	})

	if err != nil {
		return nil, fmt.Errorf("could not create chat completion: %w", err)
	}

	var objectSummary ObjectSummary
	if err := json.Unmarshal([]byte(chat.Choices[0].Message.Content), &objectSummary); err != nil {
		return nil, fmt.Errorf("could not unmarshal object summary: %w", err)
	}
	return &objectSummary, nil
}

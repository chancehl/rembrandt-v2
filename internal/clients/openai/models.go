package openai

// A struct that will be converted to a Structured Outputs response schema
type ObjectSummary struct {
	Description string `json:"description" jsonschema_description:"The plaintext description of the object in the Metropolitan Museum of Art collection (as you'd see on a placard at the museum)"`
	// NotableFacts []string `json:"notable_facts" jsonschema_description:"A few key facts about the piece"`
}

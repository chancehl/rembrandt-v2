package openai

import "github.com/invopop/jsonschema"

// Generate the JSON schema at initialization time
var objectSummaryResponseSchema = GenerateSchema[ObjectSummary]()

func GenerateSchema[T any]() interface{} {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

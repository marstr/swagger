package swagger

import "fmt"

type Scheme string

const (
	SchemeHTTP  = "http"
	SchemeHTTPS = "https"
	SchemeWS    = "ws"
	SchemeWSS   = "wss"
)

// Swagger hosts the properties belonging to an entire Swagger document.
type Swagger struct {
	Swagger             string                 `json:"swagger"`
	Info                Info                   `json:"info"`
	Host                *string                `json:"host,omitempty"`
	BasePath            *string                `json:"basePath,omitempty"`
	Schemes             *[]Scheme              `json:"schemes,omitempty"`
	Consumes            *[]string              `json:"consumes,omitempty"`
	Produces            *[]string              `json:"produces,omitempty"`
	Paths               *Paths                 `json:"paths,omitempty"`
	Definitions         *Definitions           `json:"definitions,omitempty"`
	Parameters          *map[string]Parameter  `json:"parameters,omitempty"`
	Responses           *Responses             `json:"responses,omitempty"`
	SecurityDefinitions *SecurityDefinitions   `json:"securityDefinitions,omitempty"`
	Security            *map[string][]string   `json:"security,omitempty"`
	Tags                *[]Tag                 `json:"tags,omitempty"`
	ExternalDocs        *ExternalDocumentation `json:"externalDocs,omitempty"`
}

// ResolveReference navigates a JSON reference that is rooted in this Swagger.
func (s Swagger) ResolveReference(path string) (interface{}, error) {
	current, remaining := nextJSONReferencePathElement(path)

	var decodedElement string
	if temp, err := unescapeJSONReferencePathElement(current); err == nil {
		decodedElement = temp
	} else {
		return nil, err
	}

	switch decodedElement {
	case "#":
		return s.ResolveReference(remaining)
	default:
		return nil, fmt.Errorf("element '%s' not present in Swagger", decodedElement)
	}
}

package swagger

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
	Parameters          *Parameters            `json:"parameters,omitempty"`
	Responses           *Responses             `json:"responses,omitempty"`
	SecurityDefinitions *SecurityDefinitions   `json:"securityDefinitions,omitempty"`
	Security            *map[string][]string   `json:"security,omitempty"`
	Tags                *[]Tag                 `json:"tags,omitempty"`
	ExternalDocs        *ExternalDocumentation `json:"externalDocs,omitempty"`
}

// Visit alerts a Visitor to execute against this Swagger.
func (swag Swagger) Visit(caller Visitor) {
	caller.VisitSwagger(swag)
}

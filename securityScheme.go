package swagger

type SecuritySchemeLocation string

const (
	SecuritySchemeLocationQuery  = "query"
	SecuritySchemeLocationHeader = "header"
)

type SecuritySchemeFlow string

const (
	SecuritySchemeFlowImplicit    = "implicit"
	SecuritySchemeFlowPassword    = "password"
	SecuritySchemeFlowApplication = "application"
	SecuritySchemeFlowAccessCode  = "accessCode"
)

type SecurityDefinitions map[string]SecurityScheme

type SecurityScheme struct {
	Type             string                  `json:"type"`
	Description      *string                 `json:"description,omitempty"`
	Name             *string                 `json:"name,omitempty"`
	In               *SecuritySchemeLocation `json:"in,omitempty"`
	Flow             *string                 `json:"flow,omitempty"`
	AuthorizationURL *string                 `json:"authorizationUrl,omitempty"`
	TokenURL         *string                 `json:"tokenUrl,omitempty"`
	Scopes           *OAuth2Scope            `json:"scopes,omitempty"`
}

type OAuth2Scope map[string]string

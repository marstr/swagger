package swagger

type SecuritySchemeLocation string

const (
	SecuritySchemeLocationQuery  SecuritySchemeLocation = "query"
	SecuritySchemeLocationHeader SecuritySchemeLocation = "header"
)

type SecuritySchemeFlow string

const (
	SecuritySchemeFlowImplicit    SecuritySchemeFlow = "implicit"
	SecuritySchemeFlowPassword    SecuritySchemeFlow = "password"
	SecuritySchemeFlowApplication SecuritySchemeFlow = "application"
	SecuritySchemeFlowAccessCode  SecuritySchemeFlow = "accessCode"
)

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

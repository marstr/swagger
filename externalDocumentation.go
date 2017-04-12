package swagger

type ExternalDocumentation struct {
	URL         string  `json:"url"`
	Description *string `json:"description,omitempty"`
}

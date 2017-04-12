package swagger

type Tag struct {
	Name         string                 `json:"name"`
	Description  *string                `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}

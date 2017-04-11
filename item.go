package swagger

type Items struct {
	Type             string      `json:"type"`
	Format           *string     `json:"format,omitempty"`
	Items            *Items      `json:"items,omitempty"`
	CollectionFormat *string     `json:"collectionFormat,omitempty"`
	Default          interface{} `json:"default,omitempty"`
	JSONValidation
}

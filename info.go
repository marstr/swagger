package swagger

// Info encapsulates the Info Object. http://swagger.io/specification/#info-object-17
type Info struct {
	Title          string   `json:"title"`
	Description    *string  `json:"description,omitempty"`
	TermsOfService *string  `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version"`
}

// ClearDescription removes all characters from the Description and prevents this field from being emitted.
func (info *Info) ClearDescription() {
	info.Description = nil
}

// SetDescription updates the text used to describe a Swagger document.
func (info *Info) SetDescription(value string) {
	info.Description = &value
}

// GetDescription retreives the text used to describe a Swagger document.
func (info Info) GetDescription() string {
	return *info.Description
}

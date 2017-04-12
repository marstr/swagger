package swagger

import (
	"fmt"
)

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

func (info Info) ResolveReference(path string) (interface{}, error) {
	current, remaining := nextJSONReferencePathElement(path)

	if temp, err := unescapeJSONReferencePathElement(current); err == nil {
		current = temp
	} else {
		return nil, err
	}

	switch current {
	case "":
		return info, nil
	case "contact":
		return info.Contact.ResolveReference(remaining)
	case "description":
		return info.Description, nil
	case "termsOfService":
		return info.TermsOfService, nil
	case "title":
		return info.Title, nil
	default:
		return nil, fmt.Errorf("element '%s' not present in Info", current)
	}
}

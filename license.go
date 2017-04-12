package swagger

import (
	"errors"
	"fmt"
)

// License encapsulates the License Object. Documented here: http://swagger.io/specification/#licenseObject
type License struct {
	Name string  `json:"name"`
	URL  *string `json:"url,omitempty"`
}

// ResolveReference navigates a JSON reference that is roodted in this License.
func (lic License) ResolveReference(path string) (interface{}, error) {
	current, remaining := nextJSONReferencePathElement(path)

	if remaining != "" {
		return nil, errors.New("no objects can be referenced inside of License")
	}

	if temp, err := unescapeJSONReferencePathElement(current); err == nil {
		current = temp
	} else {
		return nil, err
	}

	switch current {
	case "name":
		return lic.Name, nil
	case "url":
		return lic.URL, nil
	default:
		return nil, fmt.Errorf("element '%s' is not present in License", current)
	}
}

package swagger

import "errors"
import "fmt"

// Contact encapsutes the Contact Object. Documenation can be found here: http://swagger.io/specification/#contactObject
type Contact struct {
	Name  *string `json:"name,omitempty"`
	URL   *string `json:"url,omitempty"`
	EMail *string `json:"email,omitempty"`
}

func (contact Contact) ResolveReference(path string) (interface{}, error) {
	current, remaining := nextJSONReferencePathElement(path)

	if remaining != "" {
		return nil, errors.New("no objects can be referenced inside of Contact")
	}

	if temp, err := unescapeJSONReferencePathElement(current); err == nil {
		current = temp
	} else {
		return nil, err
	}

	switch current {
	case "":
		return contact, nil
	case "name":
		return contact.Name, nil
	case "url":
		return contact.URL, nil
	case "email":
		return contact.EMail, nil
	default:
		return nil, fmt.Errorf("unknown element '%s'", current)
	}
}

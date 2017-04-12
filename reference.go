package swagger

import "errors"
import "strings"

type ErrorReferenceNotFound struct {
	error
	Path     string
	Document Swagger
}

// ErrorMalformedReferencePath alerts of strings that don't resemble a Swagger reference.
type ErrorMalformedReferencePath struct {
	error
	Path string
}

// Referencer indicates that this element is capable of being referenced, or have children that do.
type Referencer interface {
	ResolveReference(string) (interface{}, error)
}

var (
	errorReferenceNotFoundBase      = errors.New("referenced item in swagger wasn't found")
	errorMalformedReferencePathBase = errors.New("malformed reference in swagger")
)

func newReferenceNotFoundError(path string, document Swagger) ErrorReferenceNotFound {
	return ErrorReferenceNotFound{
		error:    errorReferenceNotFoundBase,
		Path:     path,
		Document: document,
	}
}

func newReferenceMalformedReferencePath(path string) ErrorMalformedReferencePath {
	return ErrorMalformedReferencePath{
		error: errorMalformedReferencePathBase,
		Path:  path,
	}
}

func nextJSONReferencePathElement(path string) (current, remaining string) {
	index := strings.Index(path, "/")
	if index == -1 {
		current = path
	} else {
		current = path[:index]
		remaining = path[index+1:]
	}
	return
}

func escapeJSONReferencePathElement(element string) string {
	element = strings.Replace(element, "~", "~0", -1)
	element = strings.Replace(element, "/", "~1", -1)
	return element
}

func unescapeJSONReferencePathElement(element string) (string, error) {
	if strings.ContainsRune(element, '/') {
		return "", errors.New("multiple elements were passed as a single one")
	}

	// RFC6901 Section 4 states that the below must be done in this order.
	element = strings.Replace(element, "~1", "/", -1)
	element = strings.Replace(element, "~0", "~", -1)
	return element, nil
}

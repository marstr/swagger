package swagger

import "errors"

// Visitable represents
type Visitable interface {
	Visit(Visitor)
}

type Visitor interface {
	VisitDefinitions(item Definitions) error
	VisitParameters(item Parameters) error
	VisitSwagger(item Swagger) error
}

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

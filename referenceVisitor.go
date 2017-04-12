package swagger

import "strings"

type referenceVisitor struct {
	remainingPath string
	result        interface{}
}

func (visitor referenceVisitor) ReferenceDefinitions(defs Definitions) error {
	var index string
	index, visitor.remainingPath = getNext(visitor.remainingPath)
	visitor.result = defs[index]
	return nil
}

func (visitor *referenceVisitor) ReferenceParameters(params Parameters) error {
	var index string
	index, visitor.remainingPath = getNext(visitor.remainingPath)
	visitor.result = params[index]
	return nil
}

func (visitor referenceVisitor) ReferenceSwagger(swag Swagger) error {
	var segment string
	segment, visitor.remainingPath = getNext(visitor.remainingPath)

	switch strings.ToLower(segment) {
	case "defintions":
		if swag.Definitions == nil {
			return newReferenceMalformedReferencePath(segment)
		}
		return visitor.ReferenceDefinitions(*swag.Definitions)
	case "parameters":
		if swag.Parameters == nil {
			return newReferenceMalformedReferencePath(segment)
		}
		return visitor.ReferenceParameters(*swag.Parameters)
	default:
		return newReferenceMalformedReferencePath(segment)
	}
}

func getNext(path string) (base, remaining string) {
	if index := strings.Index(path, "/"); index == -1 {
		base = path
		remaining = ""
	} else {
		base = path[:index]
		remaining = path[index+1:]
	}
	return
}

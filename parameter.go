package swagger

type ParameterLocation string

const (
	ParameterLocationQuery    = "query"
	ParameterLocationHeader   = "header"
	ParameterLocationPath     = "path"
	ParameterLocationFormData = "formData"
	ParameterLocationBody     = "body"
)

type Parameters map[string]Parameter

type Parameter struct {
	Name        string            `json:"name"`
	In          ParameterLocation `json:"in"`
	Description *string           `json:"description,omitempty"`
	Required    *bool             `json:"required,omitempty"`
}

type BodyParameter struct {
	Parameter
	Schema Schema `json:"schema"`
}

type NonBodyParameter struct {
	Type            HeaderType `json:"type"`
	Format          *string
	AllowEmptyValue *bool `json:"allowEmptyValue,omitempty"`
	JSONValidation
}

// Visit alerts a Visitor
func (params Parameters) Visit(visitor Visitor) error {
	return visitor.VisitParameters(params)
}

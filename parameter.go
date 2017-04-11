package swagger

type ParameterLocation string

const (
	ParameterLocationQuery    = "query"
	ParameterLocationHeader   = "header"
	ParameterLocationPath     = "path"
	ParameterLocationFormData = "formData"
	ParameterLocationBody     = "body"
)

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
}
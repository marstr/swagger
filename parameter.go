package swagger

type ParameterLocation string

const (
	ParameterLocationQuery    ParameterLocation = "query"
	ParameterLocationHeader   ParameterLocation = "header"
	ParameterLocationPath     ParameterLocation = "path"
	ParameterLocationFormData ParameterLocation = "formData"
	ParameterLocationBody     ParameterLocation = "body"
)

// Parameter captures details about a single piece of data that will be transferred.
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

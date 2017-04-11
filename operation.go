package swagger

type Operation struct {
	Tags        *[]string `json:"tags,omitempty"`
	Summary     *string   `json:"summary,omitempty"`
	Description *string   `json:"description,omitempty"`
	OperationId *string   `json:"operationId,omitempty"`
	Consumes    *[]string `json:"consumes,omitempty"`
	Produces    *[]string `json:"produces,omitempty"`
	Deprecated  *bool     `json:"deprecated,omitempty"`
}

package swagger

// JSONValidation
type JSONValidation struct {
	Maximum          *float64  `json:"maximum,omitempty"`
	ExclusiveMaximum *bool     `json:"exclusiveMaximum,omitempty"`
	Minimum          *float64  `json:"minimum,omitempty"`
	Exclusive        *bool     `json:"exclusiveMinimum,omitempty"`
	MaxLength        *int      `json:"maxLength,omitempty"`
	MinLength        *int      `json:"minLenth,omitempty"`
	Pattern          *string   `json:"pattern,omitempty"`
	MaxItems         *int      `json:"maxItems,omitempty"`
	MinItems         *int      `json:"minItems,omitempty"`
	UniqueItem       *bool     `json:"uniqueItems,omitempty"`
	Enum             *[]string `json:"enum,omitempty"`
	MultipleOf       *float64  `json:"multipleOf,omitempty"`
}

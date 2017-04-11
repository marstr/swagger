package microsoft

import (
	"github.com/marstr/swagger"
)

type Parameter struct {
	swagger.Parameter
	SkipURL *bool `json:"x-ms-skip-url,omitempty"`
	Enum    *Enum `json:"x-ms-enum,omitempty"`
}

package swagger

// HeaderType encapsulates the `type` field of the Header Object. Documented here: http://swagger.io/specification/#header-object-68
type HeaderType string

// HeaderType known values are enumerated here.
const (
	HeaderTypeString  = "string"
	HeaderTypeNumber  = "number"
	HeaderTypeInteger = "integer"
	HeaderTypeBoolean = "boolean"
	HeaderTypeArray   = "array"
)

type HeaderCollectionFormat string

const (
	HeaderCollectionFormatCSV   = "csv"
	HeaderCollectionFormatSSV   = "ssv"
	HeaderCollectionFormatTSV   = "tsv"
	HeaderCollectionFormatPipes = "pipes"
)

type Headers map[string]Header

type Header struct {
	Description      *string                 `json:"description,omitempty"`
	Type             HeaderType              `json:"type"`
	Format           *string                 `json:"format,omitempty"`
	CollectionFormat *HeaderCollectionFormat `json:"collectionFormat,omitempty"`
	JSONValidation
}

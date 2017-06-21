package swagger

// HeaderType encapsulates the `type` field of the Header Object. Documented here: http://swagger.io/specification/#header-object-68
type HeaderType string

// HeaderType known values are enumerated here.
const (
	HeaderTypeString  HeaderType = "string"
	HeaderTypeNumber  HeaderType = "number"
	HeaderTypeInteger HeaderType = "integer"
	HeaderTypeBoolean HeaderType = "boolean"
	HeaderTypeArray   HeaderType = "array"
)

type HeaderCollectionFormat string

const (
	HeaderCollectionFormatCSV   HeaderCollectionFormat = "csv"
	HeaderCollectionFormatSSV   HeaderCollectionFormat = "ssv"
	HeaderCollectionFormatTSV   HeaderCollectionFormat = "tsv"
	HeaderCollectionFormatPipes HeaderCollectionFormat = "pipes"
)

type Headers map[string]Header

type Header struct {
	Description      *string                 `json:"description,omitempty"`
	Type             HeaderType              `json:"type"`
	Format           *string                 `json:"format,omitempty"`
	CollectionFormat *HeaderCollectionFormat `json:"collectionFormat,omitempty"`
	JSONValidation
}

package swagger

// License encapsulates the License Object. Documented here: http://swagger.io/specification/#licenseObject
type License struct {
	Name string  `json:"name"`
	URL  *string `json:"url,omitempty"`
}

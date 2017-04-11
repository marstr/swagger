package swagger

// Contact encapsutes the Contact Object. Documenation can be found here: http://swagger.io/specification/#contactObject
type Contact struct {
	Name  *string `json:"name,omitempty"`
	URL   *string `json:"url,omitempty"`
	EMail *string `json:"email,omitempty"`
}

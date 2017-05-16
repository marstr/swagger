package swagger

// Path contains each operation that is exposed via various HTTP Request Types.
type Path struct {
	Get        *Operation `json:"get,omitempty"`
	Put        *Operation `json:"put,omitempty"`
	Post       *Operation `json:"post,omitempty"`
	Delete     *Operation `json:"post,omitempty"`
	Options    *Operation `json:"options,omitempty"`
	Head       *Operation `json:"head,omitempty"`
	Patch      *Operation `json:"patch,omitempty"`
	Parameters *Parameter `json:"parameters,omitempty"`
}

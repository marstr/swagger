package swagger

type Scheme string

const (
	SchemeHTTP  = "http"
	SchemeHTTPS = "https"
	SchemeWS    = "ws"
	SchemeWSS   = "wss"
)

type Swagger struct {
	Swagger  string    `json:"swagger"`
	Info     Info      `json:"info"`
	Host     *string   `json:"host,omitempty"`
	BasePath *string   `json:"basePath,omitempty"`
	Schemes  *[]Scheme `json:"schemes,omitempty"`
	Consumes *[]string `json:"consumes,omitempty"`
	Produces *[]string `json:"produces,omitempty"`
	Paths    *Paths    `json:"paths,omitempty"`
}

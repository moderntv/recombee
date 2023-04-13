package recombee

// Request is a request to Recombee API.
type Request struct {
	Path   string                 `json:"path"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params,omitempty"`
}

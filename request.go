package recombee

// Generic request to recombee API. Use it only when request is not already implemented as part of package.
type Request struct {
	// Path in API to resource.
	Path string `json:"path"`
	// Http method.
	Method string `json:"method"`
	// Optional parameters.
	Params map[string]interface{} `json:"params,omitempty"`
}

package recombee

type Request struct {
	Path   string                 `json:"path"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params,omitempty"`
}

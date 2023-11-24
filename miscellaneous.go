package recombee

import (
	"net/http"
)

func ResetDatabase() Request {
	return Request{
		Path:   "/",
		Method: http.MethodDelete,
	}
}

func Batch() Request {
	return Request{
		Path:   "/batch/",
		Method: http.MethodPost,
	}
}

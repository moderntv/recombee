package recombee

import (
	"net/http"
)

// ResetDatabase completely erases all your data, including items, item properties, series,
// user database, purchases, ratings, detail views, and bookmarks.
// Make sure the request is never executed in the production environment!
// Resetting your database is irreversible.
func ResetDatabase() Request {
	return Request{
		Path:   "/",
		Method: http.MethodDelete,
	}
}

// Batch processing allows you to submit arbitrary sequence of requests within a single HTTPS request.
//
// Any type of request from the above documentation may be used in the Batch, and the Batch
// may combine different types of requests arbitrarily as well.
func Batch(requests []Request, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["requests"] = requests
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   "/batch/",
		Method: http.MethodPost,
		Params: params,
	}
}

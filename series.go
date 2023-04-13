package recombee

import (
	"fmt"
	"net/http"
)

// AddSeries creates a new series in the database.
func AddSeries(seriesId string) Request {
	return Request{
		Path:   fmt.Sprintf("/series/%s", seriesId),
		Method: http.MethodPut,
	}
}

// DeleteSeries deletes the series of the given seriesId from the database.
//
// Deleting a series will only delete assignment of items to it, not the items themselves!
func DeleteSeries(seriesId string) Request {
	return Request{
		Path:   fmt.Sprintf("/series/%s", seriesId),
		Method: http.MethodDelete,
	}
}

// ListSeries gets the list of all the series currently present in the database.
func ListSeries(opts ...RequestOption) Request {
	params := make(map[string]interface{})
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   "/series/lists",
		Method: http.MethodGet,
		Params: params,
	}
}

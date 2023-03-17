package recombee

import (
	"fmt"
	"net/http"
)

// Based on given seriesId adds new series/season
func AddSeries(seriesId string) Request {
	return Request{
		Path:   fmt.Sprintf("/series/%s", seriesId),
		Method: http.MethodPut,
	}
}

// Deletes existing seriesId from series view.
func DeleteSeries(seriesId string) Request {
	return Request{
		Path:   fmt.Sprintf("/series/%s", seriesId),
		Method: http.MethodDelete,
	}
}

// List all series when no options are present.
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

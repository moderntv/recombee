package recombee

import (
	"fmt"
	"net/http"
)

func AddSeries(seriesId string) Request {
	return Request{
		Path:   fmt.Sprintf("/series/%s", seriesId),
		Method: http.MethodPut,
	}
}

func DeleteSeries(seriesId string) Request {
	return Request{
		Path:   fmt.Sprintf("/series/%s", seriesId),
		Method: http.MethodDelete,
	}
}

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

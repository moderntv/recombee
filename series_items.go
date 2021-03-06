package recombee

import (
	"fmt"
	"net/http"
)

func InsertToSeries(seriesId string, itemType string, itemId string, time_ float64, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["itemType"] = itemType
	params["itemId"] = itemId
	params["time"] = time_
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/series/%s/items/", seriesId),
		Method: http.MethodPost,
		Params: params,
	}
}

func RemoveFromSeries(seriesId string, itemType string, itemId string, time_ float64) Request {
	params := make(map[string]interface{})
	params["itemType"] = itemType
	params["itemId"] = itemId
	params["time"] = time_

	return Request{
		Path:   fmt.Sprintf("/series/%s/items/", seriesId),
		Method: http.MethodDelete,
		Params: params,
	}
}

func ListSeriesItems(seriesId string) Request {
	return Request{
		Path:   fmt.Sprintf("/series/%s/items/", seriesId),
		Method: http.MethodGet,
	}
}

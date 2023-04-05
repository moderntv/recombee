package recombee

import (
	"fmt"
	"net/http"
)

// InsertToSeries inserts an existing item/series into a series of the given seriesId at a position determined by time.
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

// RemoveFromSeries removes an existing series item from the series.
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

// ListSeriesItems lists all the items present in the given series, sorted according to their time index values.
func ListSeriesItems(seriesId string) Request {
	return Request{
		Path:   fmt.Sprintf("/series/%s/items/", seriesId),
		Method: http.MethodGet,
	}
}

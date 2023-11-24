package recombee

import (
	"fmt"
	"net/http"
)

// SearchItems performs a full-text personalized search.
func SearchItems(userId string, searchQuery string, count int, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["searchQuery"] = searchQuery
	params["count"] = count
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/search/users/%s/items/", userId),
		Method: http.MethodGet,
		Params: params,
	}
}

func SearchItemSegments(userId string, searchQuery string, count int, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["searchQuery"] = searchQuery
	params["count"] = count
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/search/users/%s/item-segments/", userId),
		Method: http.MethodGet,
		Params: params,
	}
}

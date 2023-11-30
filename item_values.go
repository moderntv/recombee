package recombee

import (
	"fmt"
	"net/http"
)

// SetItemValues sets or updates (some) property values of the given item.
// The properties (columns) must be previously created by AddItemProperty.
func SetItemValues(itemId string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodPost,
		Params: params,
	}
}

// GetItemValues gets all the current property values of the given item.
func GetItemValues(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodGet,
	}
}

// UpdateMoreItems updates (some) property values of all the items that pass the filter.
func UpdateMoreItems(filter string, changes interface{}) Request {
	params := map[string]interface{}{
		"filter":  filter,
		"changes": changes,
	}
	return Request{
		Path:   "/more-items/",
		Method: http.MethodPost,
		Params: params,
	}
}

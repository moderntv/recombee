package recombee

import (
	"fmt"
	"net/http"
)

// AddItemValues sets/updates (some) property values of the given item. The properties (columns) must be previously
// created by AddItemProperty.
func AddItemValues(itemId string, values map[string]interface{}) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodPost,
		Params: values,
	}
}

// GetItemValues gets all the current property values of the given item.
func GetItemValues(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodGet,
	}
}

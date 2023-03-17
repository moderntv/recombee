package recombee

import (
	"fmt"
	"net/http"
)

// Assign values into item.
func AddItemValues(itemId string, values map[string]interface{}) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodPost,
		Params: values,
	}
}

// Returns all stored values from single item.
func GetItemValues(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodGet,
	}
}

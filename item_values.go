package recombee

import (
	"fmt"
	"net/http"
)

// SetItemValues sets or updates (some) property values of the given item.
// The properties (columns) must be previously created by AddItemProperty.
func SetItemValues(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodPost,
	}
}

// GetItemValues gets all the current property values of the given item.
func GetItemValues(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodGet,
	}
}

func UpdateMoreItems() Request {
	return Request{
		Path:   "/more-items/",
		Method: http.MethodPost,
	}
}

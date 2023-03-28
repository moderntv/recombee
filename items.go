package recombee

import (
	"fmt"
	"net/http"
)

// AddItem creates new item.
func AddItem(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodPut,
	}
}

// DeleteItem deletes existing item.
func DeleteItem(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodDelete,
	}
}

// ListItems lists all items when no opts specified.
func ListItems(opts ...RequestOption) Request {
	params := make(map[string]interface{})
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   "/items/lists",
		Method: http.MethodGet,
		Params: params,
	}
}

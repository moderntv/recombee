package recombee

import (
	"fmt"
	"net/http"
)

func AddItem(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodPut,
	}
}

func DeleteItem(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodDelete,
	}
}

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

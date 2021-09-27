package recombee

import (
	"fmt"
	"net/http"
)

func AddItemValues(itemId string, values map[string]interface{}) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodPost,
		Params: values,
	}
}

func GetItemValues(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodGet,
	}
}

package recombee

import (
	"fmt"
	"net/http"
)

func AddItemProperty(propertyName string, type_ string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/properties/%s", propertyName),
		Method: http.MethodPut,
		Params: map[string]interface{}{
			"type": type_,
		},
	}
}

func DeleteItemProperty(propertyName string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/properties/%s", propertyName),
		Method: http.MethodDelete,
	}
}

func GetItemPropertyInfo(propertyName string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/properties/%s", propertyName),
		Method: http.MethodGet,
	}
}

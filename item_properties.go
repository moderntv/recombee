package recombee

import (
	"fmt"
	"net/http"
)

// Add property (column) into items view.
func AddItemProperty(propertyName string, type_ string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/properties/%s", propertyName),
		Method: http.MethodPut,
		Params: map[string]interface{}{
			"type": type_,
		},
	}
}

// Delete property (column) from items view.
func DeleteItemProperty(propertyName string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/properties/%s", propertyName),
		Method: http.MethodDelete,
	}
}

// List all properties that are possible to set for items.
func GetItemPropertyInfo(propertyName string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/properties/%s", propertyName),
		Method: http.MethodGet,
	}
}

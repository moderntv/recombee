package recombee

import (
	"fmt"
	"net/http"
)

// AddItemProperty adds item property.
//
// Adding an item property is somehow equivalent to adding a column to the table of items. The items may be
// characterized by various properties of different types.
func AddItemProperty(propertyName string, typ string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/properties/%s", propertyName),
		Method: http.MethodPut,
		Params: map[string]interface{}{
			"type": typ,
		},
	}
}

// DeleteItemProperty deletes item property.
//
// Deleting an item property is roughly equivalent to removing a column from the table of items.
func DeleteItemProperty(propertyName string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/properties/%s", propertyName),
		Method: http.MethodDelete,
	}
}

// GetItemPropertyInfo gets information about specified item property.
func GetItemPropertyInfo(propertyName string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/properties/%s", propertyName),
		Method: http.MethodGet,
	}
}

// ListItemProperties gets the list of all the item properties in your database.
func ListItemProperties() Request {
	return Request{
		Path:   "/items/properties/list/",
		Method: http.MethodGet,
	}
}

package recombee

import (
	"fmt"
	"net/http"
)

// AddItem adds new item of the given itemId to the items catalog.
//
// All the item properties for the newly created items are set to null.
func AddItem(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodPut,
	}
}

// DeleteItem deletes an item of the given itemId from the catalog.
//
// If there are any purchases, ratings, bookmarks, cart additions, or detail views of the item present in the database,
// they will be deleted in cascade as well. Also, if the item is present in some series, it will be removed from all the
// series where present.
func DeleteItem(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s", itemId),
		Method: http.MethodDelete,
	}
}

// ListItems gets a list of IDs of items currently present in the catalog.
//
// API calls limit: 100 requests per minute. This limit can be increased for a database by the Recombee support.
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

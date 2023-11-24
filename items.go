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

// DeleteMoreItems deletes all the items that pass the filter.
//
// If an item becomes obsolete/no longer available, it is meaningful to keep it in the catalog
// (along with all the interaction data, which are very useful) and only exclude the item from recommendations.
// In such a case, use ReQL filter instead of deleting the item completely.
func DeleteMoreItems() Request {
	return Request{
		Path:   "/more-items/",
		Method: http.MethodDelete,
	}
}

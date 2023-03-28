package recombee

import (
	"fmt"
	"net/http"
)

// AddDetailView adds information into item that user has shown the detail of an item.
func AddDetailView(userId string, itemId string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["userId"] = userId
	params["itemId"] = itemId
	for _, o := range opts {
		o(params)
	}
	return Request{
		Path:   "/detailviews/",
		Method: http.MethodPost,
		Params: params,
	}
}

// DeleteDetailView removes information from item that user has shown the detail of an item.
func DeleteDetailView(userId string, itemId string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["userId"] = userId
	params["itemId"] = itemId
	for _, o := range opts {
		o(params)
	}
	return Request{
		Path:   "/detailviews/",
		Method: http.MethodDelete,
		Params: params,
	}
}

// ListItemDetailViews lists all opened details on item.
func ListItemDetailViews(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/detailviews/", itemId),
		Method: http.MethodGet,
	}
}

// ListUserDetailViews lists all opened details by user.
func ListUserDetailViews(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/detailviews/", userId),
		Method: http.MethodGet,
	}
}

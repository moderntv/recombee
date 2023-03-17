package recombee

import (
	"fmt"
	"net/http"
)

// Adds information into item that user has shown detail of item.
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

// Removes information from item that user has shown detail of item.
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

// Lists all opened details on item.
func ListItemDetailViews(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/detailviews/", itemId),
		Method: http.MethodGet,
	}
}

// Lists all opened details by user.
func ListUserDetailViews(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/detailviews/", userId),
		Method: http.MethodGet,
	}
}

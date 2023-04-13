package recombee

import (
	"fmt"
	"net/http"
)

// AddDetailView adds a detail view of the given item made by the given user.
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

// DeleteDetailView deletes an existing detail view uniquely specified by (userId, itemId, and timestamp) or all the
// detail views with the given userId and itemId if timestamp is omitted.
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

// ListItemDetailViews lists all the detail views of the given item ever made by different users.
func ListItemDetailViews(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/detailviews/", itemId),
		Method: http.MethodGet,
	}
}

// ListUserDetailViews lists all the detail views of different items ever made by the given user.
func ListUserDetailViews(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/detailviews/", userId),
		Method: http.MethodGet,
	}
}

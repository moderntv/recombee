package recombee

import (
	"fmt"
	"net/http"
)

// Adds purchase of item performed by user.
func AddPurchase(userId string, itemId string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["userId"] = userId
	params["itemId"] = itemId
	for _, o := range opts {
		o(params)
	}
	return Request{
		Path:   "/purchases/",
		Method: http.MethodPost,
		Params: params,
	}
}

// Deletes purchase of item that was previously created by user.
func DeletePurchase(userId string, itemId string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["userId"] = userId
	params["itemId"] = itemId
	for _, o := range opts {
		o(params)
	}
	return Request{
		Path:   "/purchases/",
		Method: http.MethodDelete,
		Params: params,
	}
}

// List all purchases of item.
func ListItemPurchases(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/purchases/", itemId),
		Method: http.MethodGet,
	}
}

// List all purchases of given user.
func ListUserPurchases(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/purchases/", userId),
		Method: http.MethodGet,
	}
}

package recombee

import (
	"fmt"
	"net/http"
)

// AddPurchase adds a purchase of the given item made by the given user.
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

// DeletePurchase deletes an existing purchase uniquely specified by userId, itemId, and timestamp or all the purchases
// with the given userId and itemId if timestamp is omitted.
//
// API calls limit: 1000 requests per minute. This limit can be increased for a database by the Recombee support.
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

// ListItemPurchases lists all the ever-made purchases of the given item.
//
// API calls limit: 60 requests per minute. This limit can be increased for a database by the Recombee support.
func ListItemPurchases(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/purchases/", itemId),
		Method: http.MethodGet,
	}
}

// ListUserPurchases lists all the purchases ever made by the given user.
//
// API calls limit: 60 requests per minute. This limit can be increased for a database by the Recombee support.
func ListUserPurchases(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/purchases/", userId),
		Method: http.MethodGet,
	}
}

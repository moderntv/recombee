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
func ListItemPurchases(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/purchases/", itemId),
		Method: http.MethodGet,
	}
}

// ListUserPurchases lists all the purchases ever made by the given user.
func ListUserPurchases(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/purchases/", userId),
		Method: http.MethodGet,
	}
}

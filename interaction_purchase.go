package recombee

import (
	"fmt"
	"net/http"
)

// AddPurchase adds purchase of item performed by user.
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

// DeletePurchase deletes purchase of item that was previously created by user.
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

// ListItemPurchases list all purchases of item.
func ListItemPurchases(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/purchases/", itemId),
		Method: http.MethodGet,
	}
}

// ListUserPurchases list all purchases of given user.
func ListUserPurchases(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/purchases/", userId),
		Method: http.MethodGet,
	}
}

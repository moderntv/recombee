package recombee

import (
	"fmt"
	"net/http"
)

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

func ListItemPurchases(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/purchases/", itemId),
		Method: http.MethodGet,
	}
}
func ListUserPurchases(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/purchases/", userId),
		Method: http.MethodGet,
	}
}

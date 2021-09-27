package recombee

import (
	"fmt"
	"net/http"
)

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

func ListItemDetailViews(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/detailviews/", itemId),
		Method: http.MethodGet,
	}
}

func ListUserDetailViews(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/detailviews/", userId),
		Method: http.MethodGet,
	}
}

package recombee

import (
	"fmt"
	"net/http"
)

func SetViewPortion(userId string, itemId string, portion float64, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["userId"] = userId
	params["itemId"] = itemId
	params["portion"] = portion
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   "/viewportions/",
		Method: http.MethodPost,
		Params: params,
	}
}

func DeleteViewPortion(userId string, itemId string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["userId"] = userId
	params["itemId"] = itemId
	for _, o := range opts {
		o(params)
	}
	return Request{
		Path:   "/viewportions/",
		Method: http.MethodDelete,
		Params: params,
	}
}

func ListItemViewPortions(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/viewportions/", itemId),
		Method: http.MethodGet,
	}
}

func ListUserViewPortions(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/viewportions/", userId),
		Method: http.MethodGet,
	}
}

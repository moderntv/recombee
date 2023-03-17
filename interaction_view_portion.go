package recombee

import (
	"fmt"
	"net/http"
)

// Sets how much user have watched single item.
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

// Deletes watched portion from user's single item.
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

// Get all stored view portion for item.
func ListItemViewPortions(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/viewportions/", itemId),
		Method: http.MethodGet,
	}
}

// Get all user's view portions.
func ListUserViewPortions(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/viewportions/", userId),
		Method: http.MethodGet,
	}
}

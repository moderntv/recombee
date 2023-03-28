package recombee

import (
	"fmt"
	"net/http"
)

// SetViewPortion sets how much user have watched single item. Portion is scaled from 0.0 to 1.0 (0 - 100%)
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

// DeleteViewPortion deletes watched portion from user's single item.
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

// ListItemViewPortions gets all stored view portion for an item.
func ListItemViewPortions(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/viewportions/", itemId),
		Method: http.MethodGet,
	}
}

// ListUserViewPortions gets all user items view portions.
func ListUserViewPortions(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/viewportions/", userId),
		Method: http.MethodGet,
	}
}

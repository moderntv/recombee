package recombee

import (
	"fmt"
	"net/http"
)

// AddCartAddition adds a cart addition of the given item made by the given user.
func AddCartAddition(userId string, itemId string, opts ...RequestOption) Request {
	params := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}
	for _, option := range opts {
		option(params)
	}

	return Request{
		Path:   "/cartadditions/",
		Method: http.MethodPost,
		Params: params,
	}
}

// DeleteCartAddition deletes an existing cart addition uniquely specified by userId, itemId, and timestamp
// or all the cart additions with the given userId and itemId if timestamp is omitted.
func DeleteCartAddition(userId string, itemId string, opts ...RequestOption) Request {
	params := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}
	for _, option := range opts {
		option(params)
	}

	return Request{
		Path:   "/cartadditions/",
		Method: http.MethodDelete,
		Params: params,
	}
}

// ListItemCartAddition lists all the ever-made cart additions of the given item.
func ListItemCartAddition(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/cartadditions/", itemId),
		Method: http.MethodGet,
	}
}

// ListUserCartAddition lists all the cart additions ever made by the given user.
func ListUserCartAddition(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/cartadditions/", userId),
		Method: http.MethodGet,
	}
}

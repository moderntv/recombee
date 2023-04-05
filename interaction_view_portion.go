package recombee

import (
	"fmt"
	"net/http"
)

// SetViewPortion sets viewed portion of an item (for example a video or article) by a user (at a session). If you send
// a new request with the same (userId, itemId, sessionId), the portion gets updated.
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

// DeleteViewPortion deletes an existing view portion specified by (userId, itemId, sessionId) from the database.
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

// ListItemViewPortions lists all the view portions of an item ever submitted by different users.
func ListItemViewPortions(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/viewportions/", itemId),
		Method: http.MethodGet,
	}
}

// ListUserViewPortions lists all the view portions ever submitted by the given user.
func ListUserViewPortions(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/viewportions/", userId),
		Method: http.MethodGet,
	}
}

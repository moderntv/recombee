package recombee

import (
	"fmt"
	"net/http"
)

// AddRating adds a rating of the given item made by the given user.
func AddRating(userId string, itemId string, rating float64, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["userId"] = userId
	params["itemId"] = itemId
	params["rating"] = rating
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   "/ratings/",
		Method: http.MethodPost,
		Params: params,
	}
}

// DeleteRating deletes an existing rating specified by (userId, itemId, timestamp) from the database or all the
// ratings with the given userId and itemId if timestamp is omitted.
func DeleteRating(userId string, itemId string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["userId"] = userId
	params["itemId"] = itemId
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   "/ratings/",
		Method: http.MethodDelete,
		Params: params,
	}
}

// ListItemRatings lists all the ratings of an item ever submitted by different users.
func ListItemRatings(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/ratings/", itemId),
		Method: http.MethodGet,
	}
}

// ListUserRatings lists all the ratings ever submitted by the given user.
func ListUserRatings(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/ratings/", userId),
		Method: http.MethodGet,
	}
}

package recombee

import (
	"fmt"
	"net/http"
)

// AddRating adds rating into item. Rating is scaled from -1.0 to 1.0.
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

// DeleteRating deletes specified user rating based by given itemId.
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

// LiastItemRatings lists all ratings of single item.
func ListItemRatings(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/ratings/", itemId),
		Method: http.MethodGet,
	}
}

// ListUserRatings lists all user's ratings. Returns list of items that user has ever rated.
func ListUserRatings(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/ratings/", userId),
		Method: http.MethodGet,
	}
}

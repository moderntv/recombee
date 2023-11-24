package recombee

import (
	"fmt"
	"net/http"
)

// RecommendUsersToUser gets users similar to the given user, based on the user’s past interactions
// (purchases, ratings, etc.) and values of properties.
//
// It is also possible to use POST HTTP method (for example in the case
// of a very long ReQL filter) - query parameters then become body parameters.
//
// The returned users are sorted by similarity (the first user being the most similar).
func RecommendUsersToUser(userId string, count int, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["count"] = count
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/recomms/users/%s/users/", userId),
		Method: http.MethodGet,
		Params: params,
	}
}

// RecommendUsersToItem recommends users that are likely to be interested in the given item.
//
// It is also possible to use POST HTTP method (for example in the case
// of a very long ReQL filter) - query parameters then become body parameters.
//
// The returned users are sorted by predicted interest in the item (the first user being the most interested).
func RecommendUsersToItem(itemId string, count int, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["count"] = count
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/recomms/items/%s/users/", itemId),
		Method: http.MethodGet,
		Params: params,
	}
}

package recombee

import (
	"fmt"
	"net/http"
)

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

func ListItemRatings(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/ratings/", itemId),
		Method: http.MethodGet,
	}
}
func ListUserRatings(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/ratings/", userId),
		Method: http.MethodGet,
	}
}

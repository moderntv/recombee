package recombee

import (
	"fmt"
	"net/http"
)

// SetUserValues sets or updates (some) property values of the given user.
// The properties (columns) must be previously created by Add user property.
func SetUserValues(userId string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/users/%s", userId),
		Method: http.MethodPost,
		Params: params,
	}
}

// GetUserValues gets all the current property values of the given user.
func GetUserValues(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s", userId),
		Method: http.MethodGet,
	}
}

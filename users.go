package recombee

import (
	"fmt"
	"net/http"
)

func AddUser(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s", userId),
		Method: http.MethodPut,
	}
}
func DeleteUser(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s", userId),
		Method: http.MethodDelete,
	}
}

func MergeUsers(targetUserId string, sourceUserId string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	for _, o := range opts {
		o(params)
	}
	return Request{
		Path:   fmt.Sprintf("/users/%s/merge/%s", targetUserId, sourceUserId),
		Method: http.MethodPut,
		Params: params,
	}
}

func ListUsers(opts ...RequestOption) Request {
	params := make(map[string]interface{})
	for _, o := range opts {
		o(params)
	}
	return Request{
		Path:   "/users/list",
		Method: http.MethodGet,
		Params: params,
	}
}

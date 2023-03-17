package recombee

import (
	"fmt"
	"net/http"
)

// Add new user to be recommended for.
func AddUser(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s", userId),
		Method: http.MethodPut,
	}
}

// Deletes user.
func DeleteUser(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s", userId),
		Method: http.MethodDelete,
	}
}

// Merges user interactions of two different users together. First argument is the userId where it will be stored after successful merge.
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

// List all users when no opts specified.
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

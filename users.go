package recombee

import (
	"fmt"
	"net/http"
)

// AddUser adds a new user to the database.
func AddUser(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s", userId),
		Method: http.MethodPut,
	}
}

// DeleteUser deletes a user of the given userId from the database.
//
// If there are any purchases, ratings, bookmarks, cart additions or detail views made by the user present in the
// database, they will be deleted in cascade as well.
func DeleteUser(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s", userId),
		Method: http.MethodDelete,
	}
}

// MergeUsers merges interactions (purchases, ratings, bookmarks, detail views …) of two different users under a single
// user ID. This is especially useful for online e-commerce applications working with anonymous users identified by
// unique tokens such as the session ID. In such applications, it may often happen that a user owns a persistent
// account, yet accesses the system anonymously while, e.g., putting items into a shopping cart. At some point in time,
// such as when the user wishes to confirm the purchase, (s)he logs into the system using his/her username and password.
// The interactions made under anonymous session ID then become connected with the persistent account, and merging these
// two becomes desirable.
//
// Merging happens between two users referred to as the target and the source. After the merge, all the interactions of
// the source user are attributed to the target user, and the source user is deleted.
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

// ListUsers gets a list of IDs of users currently present in the catalog.
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

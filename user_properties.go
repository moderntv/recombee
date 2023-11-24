package recombee

import (
	"fmt"
	"net/http"
)

// AddUserProperty is somehow equivalent to adding a column to the table of users.
// The users may be characterized by various properties of different types.
func AddUserProperty(propertyName string, typ string) Request {
	params := make(map[string]interface{})
	params["type"] = typ
	return Request{
		Path:   fmt.Sprintf("/users/properties/%s", propertyName),
		Method: http.MethodPut,
		Params: params,
	}
}

// DeleteUserProperty is roughly equivalent to removing a column from the table of users.
func DeleteUserProperty(propertyName string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/properties/%s", propertyName),
		Method: http.MethodDelete,
	}
}

// GetUserPropertyInfo gets information about specified user property.
func GetUserPropertyInfo(propertyName string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/properties/%s", propertyName),
		Method: http.MethodGet,
	}
}

// ListUserProperties gets the list of all the user properties in your database.
func ListUserProperties() Request {
	return Request{
		Path:   "/users/properties/list/",
		Method: http.MethodGet,
	}
}

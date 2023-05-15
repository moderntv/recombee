package recombee

import (
	"fmt"
	"net/http"
)

// AddBookmark adds a bookmark of the given item made by the given user.
func AddBookmark(userId string, itemId string, opts ...RequestOption) Request {
	params := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}
	for _, option := range opts {
		option(params)
	}

	return Request{
		Path:   "/bookmarks/",
		Method: http.MethodPost,
		Params: params,
	}
}

// DeleteBookmark deletes a bookmark uniquely specified by userId, itemId, and timestamp or
// all the bookmarks with the given `userId` and `itemId` if `timestamp` is omitted.
func DeleteBookmark(userId string, itemId string, opts ...RequestOption) Request {
	params := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}
	for _, option := range opts {
		option(params)
	}

	return Request{
		Path:   "/bookmarks/",
		Method: http.MethodDelete,
		Params: params,
	}
}

// ListItemBookmarks lists all the ever-made bookmarks of the given item.
func ListItemBookmarks(itemId string) Request {
	return Request{
		Path:   fmt.Sprintf("/items/%s/bookmarks/", itemId),
		Method: http.MethodGet,
	}
}

// ListUserBookmarks lists all the bookmarks ever made by the given user.
func ListUserBookmarks(userId string) Request {
	return Request{
		Path:   fmt.Sprintf("/users/%s/bookmarks/", userId),
		Method: http.MethodGet,
	}
}

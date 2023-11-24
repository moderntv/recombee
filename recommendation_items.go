package recombee

import (
	"fmt"
	"net/http"
)

// RecommendItemsToUser recommends top-N items that are most likely to be of high value for the given user.
//
// The most typical use cases are recommendations on the homepage, in some “Picked just for you” section, or in email.
//
// The returned items are sorted by relevance (the first item being the most relevant).
//
// Besides the recommended items, also a unique recommId is returned in the response. It can be used to:
//
//   - Let Recombee know that this recommendation was successful (e.g., user clicked one of the recommended items).
//
//   - Get subsequent recommended items when the user scrolls down (infinite scroll) or goes to the next page.
func RecommendItemsToUser(userId string, count int, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["count"] = count
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/recomms/users/%s/items/", userId),
		Method: http.MethodGet,
		Params: params,
	}
}

// RecommendItemsToItem recommends a set of items that are somehow related to one given item, X. A typical scenario is
// when the user A is viewing X. Then you may display items to the user that he might also be interested in. Recommend
// items to item request gives you Top-N such items, optionally taking the target user A into account.
//
// The returned items are sorted by relevance (the first item being the most relevant).
//
// Besides the recommended items, also a unique recommId is returned in the response. It can be used to:
//
//   - Let Recombee know that this recommendation was successful (e.g., user clicked one of the recommended items). See Reported metrics.
//
//   - Get subsequent recommended items when the user scrolls down (infinite scroll) or goes to the next page. See Recommend Next Items.
func RecommendItemsToItem(itemId string, targetUserId string, count int, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["count"] = count
	params["targetUserId"] = targetUserId
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/recomms/items/%s/items/", itemId),
		Method: http.MethodGet,
		Params: params,
	}
}

// RecommendNextItems returns items that shall be shown to a user as next recommendations when the
// user e.g. scrolls the page down (infinite scroll) or goes to the next page.
//
// It accepts recommId of a base recommendation request (e.g., request from the first page)
// and the number of items that shall be returned (count).
func RecommendNextItems(recommId string, count int, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["count"] = count
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/recomms/next/items/%s", recommId),
		Method: http.MethodGet,
		Params: params,
	}
}

// RecommendItemSegmentsToUser recommends the top Segments from a Segmentation
// for a particular user, based on the user’s past interactions.
//
// Based on the used Segmentation, this endpoint can be used for example for:
//   - Recommending the top categories for the user
//   - Recommending the top genres for the user
//   - Recommending the top brands for the user
//   - Recommending the top artists for the user
func RecommendItemSegmentsToUser(userId string, count int, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["count"] = count
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/recomms/users/%s/item-segments/", userId),
		Method: http.MethodGet,
		Params: params,
	}
}

// RecommendItemSegmentsToItem recommends Segments from a Segmentation that are the most relevant to a particular item.
//
// Based on the used Segmentation, this endpoint can be used for example for:
//   - Recommending the related categories
//   - Recommending the related genres
//   - Recommending the related brands
//   - Recommending the related artists
func RecommendItemSegmentsToItem(itemId string, targetUserId string, count int, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["count"] = count
	params["targetUserId"] = targetUserId
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/recomms/items/%s/item-segments/", itemId),
		Method: http.MethodGet,
		Params: params,
	}
}

// RecommendItemSegmentsToItemSegment recommends Segments from a result Segmentation that are the most relevant to a particular Segment from a context Segmentation.
//
// Based on the used Segmentations, this endpoint can be used for example for:
//   - Recommending the related brands to particular brand
//   - Recommending the related brands to particular category
//   - Recommending the related artists to a particular genre (assuming songs are the Items)
func RecommendItemSegmentsToItemSegment(
	contextSegmentId string,
	targetUserId string,
	count int,
	opts ...RequestOption,
) Request {
	params := make(map[string]interface{})
	params["count"] = count
	params["targetUserId"] = targetUserId
	params["contextSegmentId"] = contextSegmentId
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   "/recomms/item-segments/item-segments/",
		Method: http.MethodGet,
		Params: params,
	}
}

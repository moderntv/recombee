package recombee

import (
	"fmt"
	"net/http"
)

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

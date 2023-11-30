package recombee

import (
	"fmt"
	"net/http"
)

// CreateManualReQLSegmentation segments the items using multiple ReQL filters.
//
// Use the Add Manual ReQL Items Segment endpoint to create the individual segments.
func CreateManualReQLSegmentation(segmentationId, sourceType string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["sourceType"] = sourceType
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/segmentations/manual-reql/%s", segmentationId),
		Method: http.MethodPut,
		Params: params,
	}
}

// UpdateManualReQLSegmentation updates an existing Segmentation.
func UpdateManualReQLSegmentation(segmentationId string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/segmentations/manual-reql/%s", segmentationId),
		Method: http.MethodPost,
		Params: params,
	}
}

// AddManualReQLSegment adds a new Segment into a Manual ReQL Segmentation.
//
// The new Segment is defined by a ReQL filter that returns true for an item in case that this item belongs to the segment.
func AddManualReQLSegment(segmentationId, segmentId, filter string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["filter"] = filter
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/segmentations/manual-reql/%s/segments/%s", segmentationId, segmentId),
		Method: http.MethodPut,
		Params: params,
	}
}

// UpdateManualReQLSegment update definition of the Segment.
func UpdateManualReQLSegment(segmentationId, segmentId, filter string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["filter"] = filter
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/segmentations/manual-reql/%s/segments/%s", segmentationId, segmentId),
		Method: http.MethodPost,
		Params: params,
	}
}

// DeleteManualReQLSegment deletes a Segment from a Manual ReQL Segmentation.
func DeleteManualReQLSegment(segmentationId, segmentId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/manual-reql/%s/segments/%s", segmentationId, segmentId),
		Method: http.MethodDelete,
	}
}

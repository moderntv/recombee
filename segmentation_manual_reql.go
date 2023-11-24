package recombee

import (
	"fmt"
	"net/http"
)

// CreateManualReQLSegmentation segments the items using multiple ReQL filters.
//
// Use the Add Manual ReQL Items Segment endpoint to create the individual segments.
func CreateManualReQLSegmentation(segmentationId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/manual-reql/%s", segmentationId),
		Method: http.MethodPut,
	}
}

// UpdateManualReQLSegmentation updates an existing Segmentation.
func UpdateManualReQLSegmentation(segmentationId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/manual-reql/%s", segmentationId),
		Method: http.MethodPost,
	}
}

// AddManualReQLSegment adds a new Segment into a Manual ReQL Segmentation.
//
// The new Segment is defined by a ReQL filter that returns true for an item in case that this item belongs to the segment.
func AddManualReQLSegment(segmentationId string, segmentId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/manual-reql/%s/segments/%s", segmentationId, segmentId),
		Method: http.MethodPut,
	}
}

// UpdateManualReQLSegment update definition of the Segment.
func UpdateManualReQLSegment(segmentationId string, segmentId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/manual-reql/%s/segments/%s", segmentationId, segmentId),
		Method: http.MethodPost,
	}
}

// DeleteManualReQLSegment deletes a Segment from a Manual ReQL Segmentation.
func DeleteManualReQLSegment(segmentationId string, segmentId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/manual-reql/%s/segments/%s", segmentationId, segmentId),
		Method: http.MethodDelete,
	}
}

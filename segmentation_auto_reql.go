package recombee

import (
	"fmt"
	"net/http"
)

// CreateAutoReQLSegmentation Segment the items using a ReQL expression.
//
// For each item, the expression should return a set that contains IDs of segments to which the item belongs to.
func CreateAutoReQLSegmentation(segmentationId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/auto-reql/%s", segmentationId),
		Method: http.MethodPut,
	}
}

// UpdateAutoReQLSegmentation updates an existing Segmentation.
func UpdateAutoReQLSegmentation(segmentationId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/auto-reql/%s", segmentationId),
		Method: http.MethodPost,
	}
}

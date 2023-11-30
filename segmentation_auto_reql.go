package recombee

import (
	"fmt"
	"net/http"
)

// CreateAutoReQLSegmentation Segment the items using a ReQL expression.
//
// For each item, the expression should return a set that contains IDs of segments to which the item belongs to.
func CreateAutoReQLSegmentation(segmentationId, sourceType, expression string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["sourceType"] = sourceType
	params["expression"] = expression
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/segmentations/auto-reql/%s", segmentationId),
		Method: http.MethodPut,
		Params: params,
	}
}

// UpdateAutoReQLSegmentation updates an existing Segmentation.
func UpdateAutoReQLSegmentation(segmentationId string, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/segmentations/auto-reql/%s", segmentationId),
		Method: http.MethodPost,
		Params: params,
	}
}

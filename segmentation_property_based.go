package recombee

import (
	"fmt"
	"net/http"
)

// CreatePropertyBasedSegmentation creates a Segmentation that splits the items into segments based on values of a particular item property.
//
// A segment is created for each unique value of the property.
// In case of set properties, a segment is created for each value in the set. Item belongs to all these segments.
func CreatePropertyBasedSegmentation(segmentationId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/property-based/%s", segmentationId),
		Method: http.MethodPut,
	}
}

// UpdatePropertyBasedSegmentation updates a Property Based Segmentation
func UpdatePropertyBasedSegmentation(segmentationId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/property-based/%s", segmentationId),
		Method: http.MethodPost,
	}
}

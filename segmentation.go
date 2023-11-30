package recombee

import (
	"fmt"
	"net/http"
)

// ListSegmentations return all existing items Segmentations.
func ListSegmentations(sourceType string) Request {
	params := map[string]interface{}{
		"sourceType": sourceType,
	}
	return Request{
		Path:   "/segmentations/list/",
		Method: http.MethodGet,
		Params: params,
	}
}

// GetSegmentation gets existing Segmentation.
func GetSegmentation(segmentationId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/list/%s", segmentationId),
		Method: http.MethodGet,
	}
}

// DeleteSegmentation deletes existing Segmentation.
func DeleteSegmentation(segmentationId string) Request {
	return Request{
		Path:   fmt.Sprintf("/segmentations/list/%s", segmentationId),
		Method: http.MethodDelete,
	}
}

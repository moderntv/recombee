package recombee

import (
	"fmt"
	"net/http"
)

type Recommendations struct {
	RecommID string `json:"recommId"`
	Recomms  []struct {
		ID string `json:"id"`
	} `json:"recomms"`
	NumberNextRecommsCalls int `json:"numberNextRecommsCalls"`
}

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

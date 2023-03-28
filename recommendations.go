package recombee

import (
	"fmt"
	"net/http"
)

// Generics response from recombee API.
type Recommendations struct {
	// Unique for recommendation call. Primarilly used to send back in future requests.
	RecommID string `json:"recommId"`
	// Recommendations indentifiers that were returned as recommended content.
	Recomms []struct {
		// Point to recommended item.
		ID string `json:"id"`
	} `json:"recomms"`
	// Some of requests may desire mulitple calls to finish the recommendation process.
	NumberNextRecommsCalls int `json:"numberNextRecommsCalls"`
}

// Scenario in which the usedId is recommend count number of items.
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

// Scenario known as simillar items to user given item.
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

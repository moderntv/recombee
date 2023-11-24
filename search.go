package recombee

import (
	"fmt"
	"net/http"
)

// SearchItems performs a full-text personalized search.
func SearchItems(userId string, searchQuery string, count int, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["searchQuery"] = searchQuery
	params["count"] = count
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/search/users/%s/items/", userId),
		Method: http.MethodGet,
		Params: params,
	}
}

func SearchItemSegments(userId string, searchQuery string, count int, opts ...RequestOption) Request {
	params := make(map[string]interface{})
	params["searchQuery"] = searchQuery
	params["count"] = count
	for _, o := range opts {
		o(params)
	}

	return Request{
		Path:   fmt.Sprintf("/search/users/%s/item-segments/", userId),
		Method: http.MethodGet,
		Params: params,
	}
}

// AddSearchSynonym adds a new synonym for the Search items.
//
// When the term is used in the search query, the synonym is also used for the full-text search.
// Unless oneWay=true, it works also in the opposite way (synonym -> term).
func AddSearchSynonym() Request {
	return Request{
		Path:   "/synonyms/items/",
		Method: http.MethodPost,
	}
}

// ListSearchSynonyms gives the list of synonyms defined in the database.
func ListSearchSynonyms(opts ...RequestOption) Request {
	params := make(map[string]interface{})
	for _, o := range opts {
		o(params)
	}
	return Request{
		Path:   "/synonyms/items/",
		Method: http.MethodPost,
		Params: params,
	}
}

// DeleteAllSearchSynonyms deletes all synonyms defined in the database.
func DeleteAllSearchSynonyms() Request {
	return Request{
		Path:   "/synonyms/items/",
		Method: http.MethodDelete,
	}
}

// DeleteSearchSynonym deletes synonym of the given id. This synonym is no longer taken into account in the Search items.
func DeleteSearchSynonym(id string) Request {
	return Request{
		Path:   fmt.Sprintf("/synonyms/items/%s", id),
		Method: http.MethodDelete,
	}
}

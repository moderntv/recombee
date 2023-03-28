package recombee

import (
	"time"
)

type RequestOption = func(map[string]interface{})

// WithKeyValue is used to store own key in recombee item.
func WithKeyValue(k string, v interface{}) RequestOption {
	return func(params map[string]interface{}) {
		params[k] = v
	}
}

// WithTimestamp overrides/specifies timestamp to be stored.
func WithTimestamp(ts time.Time) RequestOption {
	return WithKeyValue("timestamp", ts.Unix())
}

// WithCascadeCreate creates cascade of elements when some of them are missing.
func WithCascadeCreate() RequestOption {
	return WithKeyValue("cascadeCreate", true)
}

// WithAmount point to amount of ordered item. Used for POST Cart addition call.
func WithAmount(amount float64) RequestOption {
	return WithKeyValue("amount", amount)
}

// WithPrice specifies price of an item.
func WithPrice(price float64) RequestOption {
	return WithKeyValue("price", price)
}

// WithProfit specifies profit when acquiring item.
func WithProfit(profit float64) RequestOption {
	return WithKeyValue("profit", profit)
}

// WithRecommId can be used when received from previous API calls. Improves accuracy of recommendations.
func WithRecommId(id string) RequestOption {
	return WithKeyValue("recommId", id)
}

// WithDuration sets duration of an interaction.
func WithDuration(dur time.Duration) RequestOption {
	return WithKeyValue("duration", dur)
}

// WithSessionId specifies user session ID.
func WithSessionId(id string) RequestOption {
	return WithKeyValue("sessionId", id)
}

// WithFilter is used as a custom filter defined in recombee filters.
func WithFilter(filter string) RequestOption {
	return WithKeyValue("filter", filter)
}

// WithCount limits the number of received elements/items.
func WithCount(count int) RequestOption {
	return WithKeyValue("count", count)
}

// WithOffset offests the API received elements/items.
func WithOffset(offset int) RequestOption {
	return WithKeyValue("offset", offset)
}

// WithReturnProperties returns properties of items.
func WithReturnProperties(enable bool) RequestOption {
	return WithKeyValue("returnProperties", enable)
}

// WithIncludedProperties includes properties to items.
func WithIncludedProperties(enable bool) RequestOption {
	return WithKeyValue("includedProperties", enable)
}

// WithScenario applies a scenario that is defined in recombee. Default ones or custom ones
func WithScenario(scenario string) RequestOption {
	return WithKeyValue("scenario", scenario)
}

// WithUserImpact specifies how much impact on user item has.
func WithUserImpact(impact float64) RequestOption {
	return WithKeyValue("userImpact", impact)
}

// WithBooster boosts item to be returned more often.
func WithBooster(booster string) RequestOption {
	return WithKeyValue("booster", booster)
}

// WithLogic is typicaly string which points to recombee Models/Search/Advetings logic or
// own specific logic for given recommendation request.
func WithLogic(logic interface{}) RequestOption {
	return WithKeyValue("logic", logic)
}

// WithMinRelevance is relevance of recommended item. Possible values are "low", "medium", "high".
// Default is "low, meaning that the recombee system attempts to recommend a number of items to count at any cost.
func WithMinRelevance(min float64) RequestOption {
	return WithKeyValue("minRelevance", min)
}

// WithRotationRate is used when `targetUserId` is provided. Should be used when user recommendations are resolved in real-time.
// The question is, how much the recommendation should change? When set to 0 remains the recommendation the same.
// When set to 1 means maximal rotation of recommended items for user.
func WithRotationRate(rate float64) RequestOption {
	return WithKeyValue("rotationRate", rate)
}

// WithRotationTime is used when `targetUserId` is provided. Taking `rotationRate` into account.
// Speicfies how long it takes recommended item from previous rotation to recover from penalization.
// t is specified as seconds with precision to milliseconds.
func WithRotationTime(t float64) RequestOption {
	return WithKeyValue("rotationTime", t)
}

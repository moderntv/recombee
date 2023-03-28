package recombee

import (
	"time"
)

type RequestOption = func(map[string]interface{})

// Used to store own key in recombee item.
func WithKeyValue(k string, v interface{}) RequestOption {
	return func(params map[string]interface{}) {
		params[k] = v
	}
}

// Overrides/specifies timestamp to be stored.
func WithTimestamp(ts time.Time) RequestOption {
	return WithKeyValue("timestamp", ts.Unix())
}

// Creates cascade of elements when some of them are missing.
func WithCascadeCreate() RequestOption {
	return WithKeyValue("cascadeCreate", true)
}

// Amount of ordered item. Used for POST Cart addition.
func WithAmount(amount float64) RequestOption {
	return WithKeyValue("amount", amount)
}

// Price of item.
func WithPrice(price float64) RequestOption {
	return WithKeyValue("price", price)
}

// Profit when acquiring item.
func WithProfit(profit float64) RequestOption {
	return WithKeyValue("profit", profit)
}

// Specified when received from previous API calls. Improves accuracy of recommendations.
func WithRecommId(id string) RequestOption {
	return WithKeyValue("recommId", id)
}

// Duration of interaction.
func WithDuration(dur time.Duration) RequestOption {
	return WithKeyValue("duration", dur)
}

// Specifies user session ID.
func WithSessionId(id string) RequestOption {
	return WithKeyValue("sessionId", id)
}

// Uses custom filter defined in recombee filters.
func WithFilter(filter string) RequestOption {
	return WithKeyValue("filter", filter)
}

// Limits the number of received elements/items.
func WithCount(count int) RequestOption {
	return WithKeyValue("count", count)
}

// Offests the API received elements/items.
func WithOffset(offset int) RequestOption {
	return WithKeyValue("offset", offset)
}

// Returns properties of items.
func WithReturnProperties(enable bool) RequestOption {
	return WithKeyValue("returnProperties", enable)
}

// Includes properties to items.
func WithIncludedProperties(enable bool) RequestOption {
	return WithKeyValue("includedProperties", enable)
}

// Applies scenario that is defined in recombee.
func WithScenario(scenario string) RequestOption {
	return WithKeyValue("scenario", scenario)
}

// Specifies how much impact on user item has.
func WithUserImpact(impact float64) RequestOption {
	return WithKeyValue("userImpact", impact)
}

// Boosts item to be returned more often
func WithBooster(booster string) RequestOption {
	return WithKeyValue("booster", booster)
}

// Typicaly string which points to recombee Models/Search/Advetings logic or
// own specific logic for given recommendation request.
func WithLogic(logic interface{}) RequestOption {
	return WithKeyValue("logic", logic)
}

// Relevance of recommended item. Possible values are "low", "medium", "high".
// Default is "low, meaning that the recombee system attempts to recommend a number of items to count at any cost.
func WithMinRelevance(min float64) RequestOption {
	return WithKeyValue("minRelevance", min)
}

func WithRotationRate(rate string) RequestOption {
	return WithKeyValue("rotationRate", rate)
}

// Used when `targetUserId` is provided. Taking `rotationRate` into account.
// Speicfies how long it takes recommended item from previous rotation to recover from penalization.
func WithRotationTime(t float64) RequestOption {
	return WithKeyValue("rotationTime", t)
}

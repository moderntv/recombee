package recombee

import (
	"time"
)

type RequestOption = func(map[string]interface{})

func WithKeyValue(k string, v interface{}) RequestOption {
	return func(params map[string]interface{}) {
		params[k] = v
	}
}

func WithTimestamp(ts time.Time) RequestOption {
	return WithKeyValue("timestamp", ts.Unix())
}
func WithCascadeCreate() RequestOption {
	return WithKeyValue("cascadeCreate", true)
}
func WithAmount(amount float64) RequestOption {
	return WithKeyValue("amount", amount)
}
func WithPrice(price float64) RequestOption {
	return WithKeyValue("price", price)
}
func WithProfit(profit float64) RequestOption {
	return WithKeyValue("profit", profit)
}
func WithRecommId(id string) RequestOption {
	return WithKeyValue("recommId", id)
}
func WithDuration(dur time.Duration) RequestOption {
	return WithKeyValue("duration", dur)
}
func WithSessionId(id string) RequestOption {
	return WithKeyValue("sessionId", id)
}
func WithFilter(filter string) RequestOption {
	return WithKeyValue("filter", filter)
}
func WithCount(count int) RequestOption {
	return WithKeyValue("count", count)
}
func WithOffset(offset int) RequestOption {
	return WithKeyValue("offset", offset)
}
func WithReturnProperties(enable bool) RequestOption {
	return WithKeyValue("returnProperties", enable)
}
func WithIncludedProperties(enable bool) RequestOption {
	return WithKeyValue("includedProperties", enable)
}
func WithScenario(scenario string) RequestOption {
	return WithKeyValue("scenario", scenario)
}
func WithUserImpact(impact float64) RequestOption {
	return WithKeyValue("userImpact", impact)
}
func WithBooster(booster string) RequestOption {
	return WithKeyValue("booster", booster)
}
func WithLogic(logic interface{}) RequestOption {
	return WithKeyValue("logic", logic)
}
func WithMinRelevance(min float64) RequestOption {
	return WithKeyValue("minRelevance", min)
}
func WithRotationRate(rate float64) RequestOption {
	return WithKeyValue("rotationRate", rate)
}
func WithRotationTime(t float64) RequestOption {
	return WithKeyValue("rotationTime", t)
}

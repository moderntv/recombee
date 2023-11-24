package recombee

type Recommendations struct {
	RecommID string `json:"recommId"`
	Recomms  []struct {
		ID string `json:"id"`
	} `json:"recomms"`
	NumberNextRecommsCalls int `json:"numberNextRecommsCalls"`
}

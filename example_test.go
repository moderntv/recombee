package recombee_test

import (
	"github.com/moderntv/recombee"
)

func ExampleWithRotationTime() {
	recombee.WithRotationTime(7200.0)
	// means that items recommended less than 2 hours are penalized.
}

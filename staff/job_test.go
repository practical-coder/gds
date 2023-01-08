package staff

import (
	"fmt"
	"testing"
)

func TestJob(t *testing.T) {
	fixed := FixedPriceJob{"Stucco House", 34760.0}
	partTime := HourlyJob{"Landscaping", 40.0, 50}
	jobs := []JobInterface{fixed, partTime}
	totalCost := TotalJobCost(jobs)
	fmt.Printf("Total job cost: $%0.2f\n", totalCost)
}

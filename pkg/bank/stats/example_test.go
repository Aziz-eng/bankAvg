package stats

import (
	"fmt"

	"github.com/Aziz-eng/bankAvg/pkg/bank/types"
)

// for
func ExampleAvg() {
	payment := []types.Payment{
		{
			ID:       1,
			Amount:   100,
			Category: "Market",
		},
		{
			ID:       2,
			Amount:   400,
			Category: "Petrol",
		},
	}
	result := Avg(payment)
	fmt.Println(result)
	// Output: 250
}

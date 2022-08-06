package concepts

import (
	"fmt"
	"testing"
)

func higherInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test functions must start with Test and take in *testing.T

// TestHigherIntBasic This is how we can declare a single test
func TestHigherIntBasic(t *testing.T) {
	ans := higherInt(5, 10)

	if ans != 10 {
		t.Errorf("higherInt(5, 10) = %d. Should be 10", ans)
	}
}

// TestHigherIntTable To test multiple times, or with multiple different cases we can create a tests struct
// We can then run t.Run the same we ran our individual test above
func TestHigherIntTable(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{5, 10, 10},
		{42, 12, 42},
		{1, -4, 1},
	}

	for _, tt := range tests {

		testName := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			ans := higherInt(tt.a, tt.b)

			if ans != tt.want {
				t.Errorf("higherInt(%d, %d) = %d. Should be %d", tt.a, tt.b, ans, tt.want)
			}
		})
	}
}

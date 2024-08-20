package namespace

import (
	"math"
	"testing"
)

func TestHandle_Type(t *testing.T) {
	testData := []int{math.MinInt, -1, 0, 1, math.MaxInt}
	for i := range testData {
		_ = Handle(i)
	}
}

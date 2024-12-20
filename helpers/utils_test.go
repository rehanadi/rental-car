package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCost(t *testing.T) {
	subtotalCost, taxCost, totalCost := CalculateCost(100, 2)
	assert.Equal(t, 200.0, subtotalCost)
	assert.Equal(t, 24.0, taxCost)
	assert.Equal(t, 224.0, totalCost)
}

package calculator

import (
	"testing"
	"github.com/stretchr/testify/assert" 
)

func TestCalculate(t *testing.T) {
	result := Calculate()
	assert.NotNil(t, result)
}
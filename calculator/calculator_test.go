package calculator

import (
	"testing"
	"github.com/stretchr/testify/assert" 
)

func TestCalculate(t *testing.T) {
	result, error := Calculate("05/11/2015 17:00")
	assert.Equal(t, result, 0);
	assert.Nil(t, error);
}

func TestShouldRequireAStartTime(t *testing.T) {
	result, error := Calculate("");
	assert.Equal(t, result, 0);
	assert.NotNil(t, error);
}
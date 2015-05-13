package calculator

import (
	"testing"
	"github.com/stretchr/testify/assert" 
)

func TestCalculate(t *testing.T) {
	result, error := Calculate("05/11/2015 5:00pm", "05/11/2015 05:00pm", "05/11/2015 05:00pm")
	assert.Equal(t, result, 0);
	assert.Nil(t, error);
}

func TestShouldRequireAStartTime(t *testing.T) {
	result, error := Calculate("", "05/11/2015 05:00pm", "05/11/2015 05:00pm");
	assert.Equal(t, result, 0);
	assert.NotNil(t, error);
}

func TestShouldRequireAStartTimeToBeParsable(t *testing.T) {
	result, error := Calculate("blah", "05/11/2015 05:00pm", "05/11/2015 05:00pm");
	assert.Equal(t, result, 0);
	assert.NotNil(t, error);
}

func TestShouldRequireALeaveTime(t *testing.T) {
	result, error := Calculate("05/11/2015 05:00pm", "", "05/11/2015 05:00pm");
	assert.Equal(t, result, 0);
	assert.NotNil(t, error);
}

func TestShouldRequireABedTime(t *testing.T) {
	result, error := Calculate("05/11/2015 05:00pm", "05/11/2015 05:00pm", "");
	assert.Equal(t, result, 0);
	assert.NotNil(t, error);
}
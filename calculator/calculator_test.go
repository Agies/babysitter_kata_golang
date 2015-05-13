package calculator

import (
	"testing"
	"github.com/stretchr/testify/assert" 
)

func TestCalculate(t *testing.T) {
	_, error := Calculate("05/11/2015 5:00pm", "05/11/2015 05:00pm", "05/11/2015 05:00pm")
	assert.Nil(t, error);
}

func TestShouldRequireAStartTime(t *testing.T) {
	_, error := Calculate("", "05/11/2015 05:00pm", "05/11/2015 05:00pm");
	assert.NotNil(t, error);
}

func TestShouldRequireAStartTimeToBeParsable(t *testing.T) {
	_, error := Calculate("blah", "05/11/2015 05:00pm", "05/11/2015 05:00pm");
	assert.NotNil(t, error);
}

func TestShouldRequireALeaveTime(t *testing.T) {
	_, error := Calculate("05/11/2015 05:00pm", "", "05/11/2015 05:00pm");
	assert.NotNil(t, error);
}

func TestShouldRequireALeaveTimeToBeParsable(t *testing.T) {
	_, error := Calculate("05/11/2015 05:00pm", "blah", "05/11/2015 05:00pm");
	assert.NotNil(t, error);
}

func TestShouldRequireABedTime(t *testing.T) {
	_, error := Calculate("05/11/2015 05:00pm", "05/11/2015 05:00pm", "");
	assert.NotNil(t, error);
}

func TestShouldRequireABedTimeToBeParsable(t *testing.T) {
	_, error := Calculate("05/11/2015 05:00pm", "05/11/2015 05:00pm", "blah");
	assert.NotNil(t, error);
}

func TestShouldRequireAStartTimeOfNoEarlierThan5PM(t *testing.T) {
	_, error := Calculate("05/11/2015 04:00pm", "05/11/2015 05:00pm", "05/11/2015 05:00pm");
	assert.NotNil(t, error);
}
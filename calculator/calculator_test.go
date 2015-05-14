package calculator

import (
	"testing"
	"github.com/stretchr/testify/assert" 
)

func TestCalculate(t *testing.T) {
	_, error := Calculate("05/11/2015 5:00pm", "05/11/2015 5:00pm", "05/11/2015 5:00pm")
	assert.Nil(t, error);
}

func TestShouldRequireAStartTime(t *testing.T) {
	_, error := Calculate("", "05/11/2015 5:00pm", "05/11/2015 5:00pm");
	assert.NotNil(t, error, error);
}

func TestShouldRequireAStartTimeToBeParsable(t *testing.T) {
	_, error := Calculate("blah", "05/11/2015 5:00pm", "05/11/2015 5:00pm");
	assert.NotNil(t, error);
}

func TestShouldRequireALeaveTime(t *testing.T) {
	_, error := Calculate("05/11/2015 5:00pm", "", "05/11/2015 5:00pm");
	assert.NotNil(t, error);
}

func TestShouldRequireALeaveTimeToBeParsable(t *testing.T) {
	_, error := Calculate("05/11/2015 5:00pm", "blah", "05/11/2015 5:00pm");
	assert.NotNil(t, error);
}

func TestShouldRequireABedTime(t *testing.T) {
	_, error := Calculate("05/11/2015 5:00pm", "05/11/2015 5:00pm", "");
	assert.NotNil(t, error);
}

func TestShouldRequireABedTimeToBeParsable(t *testing.T) {
	_, error := Calculate("05/11/2015 5:00pm", "05/11/2015 5:00pm", "blah");
	assert.NotNil(t, error);
}

func TestShouldRequireAStartTimeOfNoEarlierThan5pm(t *testing.T) {
	_, error := Calculate("05/11/2015 04:00pm", "05/11/2015 5:00pm", "05/11/2015 5:00pm");
	assert.NotNil(t, error);
}

func TestShouldrequirealeavetimetobenolaterthan4amthefollowingday(t*testing.T){
     _, error := Calculate("05/11/2015 5:00pm", "05/12/2015 5:00am", "05/11/2015 8:00pm");
	 assert.NotNil(t, error);
}
func TestShouldrequireabedtimetobeparsable(t *testing.T){
	_, error := Calculate("05/11/2015 5:00pm", "05/11/2015 4:00pm", "blah");        
	assert.NotNil(t, error);
}
func TestShouldpay36beforebedfor3hours(t *testing.T){
	amount, error := Calculate("05/11/2015 5:00pm", "05/11/2015 8:00pm", "05/11/2015 8:00pm");
	assert.Equal(t, amount, 36);
	assert.Nil(t, error);
}
func TestShouldpay12hrbeforebedequaling48dollarif4hrs(t *testing.T){
	amount, error := Calculate("05/11/2015 5:00pm", "05/11/2015 9:00pm", "05/11/2015 9:00pm");
	assert.Equal(t, amount, 48);
	assert.Nil(t, error);
} 
func TestShouldpay0prebedtimeifstarttimeafterbedtime(t *testing.T){
	BeforeBedtimePayment = 12;
	amount, error := Calculate("05/11/2015 11:00pm", "05/11/2015 8:00pm", "05/11/2015 8:00pm");
	assert.Equal(t, amount, 0);
	assert.Nil(t, error);
} 
func TestShouldpay32afterbedandbeforemidnightif4hrs(t *testing.T){
	BeforeBedtimePayment = 12;
	amount, error := Calculate("05/11/2015 11:00pm", "05/12/2015 12:00am", "05/11/2015 8:00pm");
	assert.Equal(t, amount, 32);
	assert.Nil(t, error);
} 
func TestShouldpay8hrafterbedandbeforemidnight(t *testing.T){
	BeforeBedtimePayment = 12;
	amount, error := Calculate("05/11/2015 11:00pm", "05/12/2015 12:00am", "05/11/2015 7:00pm");
	assert.Equal(t, amount, 40);
	assert.Nil(t, error);
}
func TestShouldpay16aftermidnightif1hr(t *testing.T){
	BeforeBedtimePayment = 0;
	amount, error := Calculate("05/11/2015 7:00pm", "05/12/2015 1:00am", "05/12/2015 12:00am");
	assert.Equal(t, amount, 16);
	assert.Nil(t, error);
} 
func TestShouldpayonlyforwholehours(t *testing.T){
	BeforeBedtimePayment = 12;
	amount, error := Calculate("05/11/2015 5:15pm", "05/12/2015 1:45am", "05/11/2015 8:00pm");
	assert.Equal(t, amount, 72);
	assert.Nil(t, error);
}
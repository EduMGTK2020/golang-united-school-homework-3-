package homework

import "testing"

func TestAverage(t *testing.T) {
	var input = [15]float32 {1,2,3,4,5,6}
	var expAv float32 = 3.5
	if expAv != average(input){
		t.Errorf("Unexpected result:\n\tExpected: %f\n\tGot: %f", expAv, average(input))
	}
}

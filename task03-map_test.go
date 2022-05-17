package homework

import "testing"

func TestSortMapValues(t *testing.T) {
	var input = map[int]string{2: "a", 0: "b", 1: "c"}
	var got = sortMapValues(input)

	var expect = []string{"b", "c", "a"}
	var len = len(expect)

	for i := 0; i < len; i++ {
		if input[i] != got[i] {
			t.Errorf("Unexpected result:\n\tExpected: %s\n\tGot: %s", expect[i], input[i])
		}
	}
}

package homework

import "testing"

func TestReverse(t *testing.T) {
	var input = []int64{1, 2, 5, 15}
	var got = reverse(input)
	var len = len(input)

	var expRev = []int64{15, 5, 2, 1}

	for i := 0; i < len; i++ {
		if got[i] != expRev[i] {
			t.Errorf("Unexpected result:\n\tExpected: %d\n\tGot: %d", expRev[i], got[i])
		}
	}
}

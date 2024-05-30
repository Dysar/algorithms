package __169_majority_element

import "testing"

func Test_majorityElement(t *testing.T) {
	m := majorityElement([]int{2, 2, 3, 3, 2})
	if m != 2 {
		t.FailNow()
	}
}

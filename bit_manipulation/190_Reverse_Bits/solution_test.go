package main

import "testing"

func Test_reverseBits(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := reverseBits(43261596)
		if output != 964176192 {
			t.Errorf("Output was incorrect, got: %d, want: %d.", output, 964176192)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		output := reverseBits(4294967293)
		if output != 3221225471 {
			t.Errorf("Output was incorrect, got: %d, want: %d.", output, 3221225471)
		}
	})

	t.Run("test", func(t *testing.T) {
		answer, num := 0, 43261596
		t.Logf("num&1: '%b' & 1: '%b'", num, num&1)
		t.Logf("(num&1)<<(31-0): '%b' << 31: '%b'", num&1, (num&1)<<(31-0))
		t.Logf("answer|(num&1)<<(31-0): '%b' | '%b': '%b'", answer, (num&1)<<(31-0), answer|(num&1)<<(31-0))
		answer = answer | (num&1)<<(31-0)
		t.Logf("num>>1: '%b' >> 1: '%b'", num, num>>1)
		num = num >> 1

		t.Logf("num&1: '%b' & 1: '%b'", num, num&1)
		t.Logf("(num&1)<<(31-0): '%b' << 31: '%b'", num&1, (num&1)<<(31-1))
		t.Logf("answer|(num&1)<<(31-0): '%b' | '%b': '%b'", answer, (num&1)<<(31-1), answer|(num&1)<<(31-1))
		answer = answer | (num&1)<<(31-1)
		t.Logf("num>>1: '%b' >> 1: '%b'", num, num>>1)
		num = num >> 1
	})
}

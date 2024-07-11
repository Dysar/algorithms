package main

func reverseBits(num uint32) uint32 {
	var answer uint32
	for i := 0; i < 32; i++ {
		answer = answer | (num&1)<<(31-i) // shifts all bits in
		num = num >> 1                    // shifts all bits in num to the right 1 time
	}
	return answer
}

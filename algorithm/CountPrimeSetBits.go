package main

import (
	"fmt"
	"math/bits"
)

func countPrimeSetBits(left int, right int) (ans int) {
	for i := left; i <= right; i++ {
		if (1<<bits.OnesCount(uint(i)))&665772 != 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countPrimeSetBits(10, 15))
}

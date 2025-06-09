package main

import (
	"fmt"
	"math/rand/v2"
)

// countBits returns (leading zeros, zeros, ones)
func countBits(num uint64) (int, int, int) {
	var zeros, ones int
	for i := 64; i > 0; i-- {
		if num == 0 {
			// the remaining bits are all leading zeros
			return i, zeros, ones
		}
		if num&0x1 == 1 {
			ones++
		} else {
			zeros++
		}
		num = num >> 1
	}
	return 0, zeros, ones
}

func main() {
	// Count is the number of iterations we generate
	count := 100_000_000_000

	histogramLeading := [64]uint64{}
	histogramZeros := [64]uint{}
	histogramOnes := [64]uint{}

	for range count {
		val := rand.Uint64()
		leading, zeros, ones := countBits(val)
		histogramLeading[leading]++
		histogramZeros[zeros]++
		histogramOnes[ones]++
	}

	fmt.Println("T,leading,zeros,ones")

	for i := range 64 {
		fmt.Printf("%d,%d,%d,%d\n", i, histogramLeading[i], histogramZeros[i], histogramOnes[i])
	}
}

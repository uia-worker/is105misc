package main

import (
	"fmt"
	"math"
	"strings"
)

func countBitsForLargeInteger(numStr string) int {
	// Fjern eventuelle foranstående nuller
	numStr = strings.TrimLeft(numStr, "0")

	// Hvis numStr er tom etter fjerning av nuller, returner 1 (for 0)
	if numStr == "" {
		return 1
	}

	// Finn lengden av strengen og beregn antall bits
	strLen := len(numStr)
	bitCount := int(math.Ceil(float64(strLen) * 3.32193)) // Log2(10) er omtrent 3.32193

	return bitCount
}

func main() {
	// Large integer value as a string
	largeIntStr := "1234567890123456789012345678901234567890"

	// Calculate the number of bits required
	numBits := countBitsForLargeInteger(largeIntStr)

	// Print the result
	fmt.Printf("Number of bits required to represent %s: %d\n", largeIntStr, numBits)
}


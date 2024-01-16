package main

import (
	"fmt"
	"math"
)

const (
	base   = uint64(1e18) // Base for each element in the custom big integer (1e18 = 10^18)
	digits = 18           // Number of digits in each element
)

type CustomBigInt struct {
	parts []uint64 // Each part stores 'digits' number of decimal digits
}

func NewCustomBigInt(value uint64) *CustomBigInt {
	return &CustomBigInt{parts: []uint64{value}}
}

func (a *CustomBigInt) Add(b *CustomBigInt) *CustomBigInt {
	c := &CustomBigInt{}
	carry := uint64(0)

	// Perform addition part by part
	for i := 0; i < len(a.parts) || i < len(b.parts) || carry > 0; i++ {
		var aPart, bPart uint64

		if i < len(a.parts) {
			aPart = a.parts[i]
		}

		if i < len(b.parts) {
			bPart = b.parts[i]
		}

		sum := aPart + bPart + carry
		c.parts = append(c.parts, sum%base)
		carry = sum / base
	}

	return c
}

func (a *CustomBigInt) String() string {
	if len(a.parts) == 0 {
		return "0"
	}

	result := fmt.Sprintf("%d", a.parts[len(a.parts)-1])
	format := fmt.Sprintf("%%0%dd", digits)

	for i := len(a.parts) - 2; i >= 0; i-- {
		result += fmt.Sprintf(format, a.parts[i])
	}

	return result
}

func main() {
	// Create custom big integers
	num1 := NewCustomBigInt(math.MaxUint64)
	num2 := NewCustomBigInt(12345678901234567890000000000000000999882)

	// Add the custom big integers
	sum := num1.Add(num2)

	// Print the result
        fmt.Println("math.MaxUint64:", num1)
	fmt.Println("number        :", num2)
	fmt.Println("Sum:", sum)
}


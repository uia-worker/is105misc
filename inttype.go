package main

import (
	"fmt"
	"math"
	"strconv"
)

func findGoTypeAndBitSize(input string) (string, int) {
	// Prøv å konvertere inndata til forskjellige Go-typer og finn antall bits
	intVal, err := strconv.ParseInt(input, 10, 64)
	if err == nil {
		if intVal >= math.MinInt8 && intVal <= math.MaxInt8 {
			return "int8", 8
		} else if intVal >= math.MinInt16 && intVal <= math.MaxInt16 {
			return "int16", 16
		} else if intVal >= math.MinInt32 && intVal <= math.MaxInt32 {
			return "int32", 32
		} else {
			return "int64", 64
		}
	}

	uintVal, err := strconv.ParseUint(input, 10, 64)
	if err == nil {
		if uintVal <= math.MaxUint8 {
			return "uint8", 8
		} else if uintVal <= math.MaxUint16 {
			return "uint16", 16
		} else if uintVal <= math.MaxUint32 {
			return "uint32", 32
		} else {
			return "uint64", 64
		}
	}

	return "big.Int", -1 // Bruk big.Int hvis ingen av de innebygde typene fungerer
}

func main() {
	// Ta et tilfeldig stort heltall som inndata
	input := "1234567890123456789012345678901234567890"

	// Finn Go-typen og antall bits som trengs for representasjon
	goType, bitSize := findGoTypeAndBitSize(input)

	// Skriv ut resultatet
	fmt.Printf("Input-verdi: %s\n", input)
	fmt.Printf("Relevant Go-type: %s\n", goType)
	fmt.Printf("Antall bits: %d\n", bitSize)
}


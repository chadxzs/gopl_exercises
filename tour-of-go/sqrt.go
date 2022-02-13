package main

import (
	"fmt"
	"math"
    "os"
    "strconv"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prev := 0.0

	for math.Abs(z - prev) > 0.0000001 {
		prev = z
		z -= (z*z - x) / (2*z)
	}

	return z
}

func main() {
    arg, err := strconv.ParseFloat(os.Args[1], 64)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Invalid number: %s\n", os.Args[1])
        os.Exit(1)
    }
    fmt.Printf("Calculating sqrt of %.0f\n", arg)
    fmt.Printf("My result is: %.10f\n", Sqrt(arg))
	fmt.Printf("Standard library result is: %.10f\n", math.Sqrt(arg))
}


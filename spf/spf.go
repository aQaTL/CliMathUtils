package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func gcd(a, b float64) float64 {
	if a == 0 || b == 0 {
		return math.Max(a, b)
	} else {
		return gcd(math.Min(a, b), math.Mod(math.Max(a, b), math.Min(a, b)))
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ")
		fmt.Println("	spf [number_1] [number_2]")
		os.Exit(1)
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		panic(err)
	}
	gcdRes := gcd(a, b)
	fmt.Println("GCD:", gcdRes)
	fmt.Print(a, "/", b, " = ", a/gcdRes, "/", b/gcdRes, "\n")
}

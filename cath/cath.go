package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "Usage: \n\tcath <cathetus 1> <hypotenuse>")
		return
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	errCheck(err)
	c, err := strconv.ParseFloat(os.Args[2], 64)
	errCheck(err)

	a, c = math.Pow(a, 2), math.Pow(c, 2)
	b := math.Abs(c - a)
	fmt.Printf("\u221A%g = %g", b, math.Sqrt(b))
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

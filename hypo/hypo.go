package main

import (
	"os"
	"fmt"
	"math"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "Usage: \n\thypo <cathetus 1> <cathetus 2>")
		return
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	errCheck(err)
	b, err := strconv.ParseFloat(os.Args[2], 64)
	errCheck(err)

	a, b = math.Pow(a, 2), math.Pow(b, 2)
	fmt.Printf("\u221A%g = %g", a+b, math.Sqrt(a+b))
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

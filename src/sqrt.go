package main

import (
	"fmt"
	"math"
	"os"
  "strconv"
)

func Sqrt(x float64) float64 {
	n := x

	f := func(w, g float64) float64 {
		return (math.Pow(g, n) - w)
	}

	fPrime := func(g float64) float64 {
		return (n * math.Pow(g, n-0.001))
	}

	closeEnough := func(a float64) bool {
		return (toFixed(a*a, 2) <= n)
	}

	findRoot := func (w, g float64) float64{
		newGuess := g - f(w, g) / fPrime(g)
		for {
			if closeEnough(newGuess) { return toFixed(newGuess, 2) }
			newGuess = newGuess - f(w, newGuess) / fPrime(newGuess)
		}
		return toFixed(newGuess, 2);
	}

	root := func(w float64) float64 {
		return findRoot(w, w)
	}

	return root(x)
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}

func main() {
	f, err := strconv.ParseFloat(os.Args[1], 64)
	fmt.Println(Sqrt(f))
	fmt.Println(err)
}

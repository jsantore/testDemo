package main

import (
	"fmt"
	"math"
)

func calcDistance(x1, y1, x2, y2 float64) float64 {
	fmt.Println(x1, x2, y1, y2)
	diff1 := x1 - x2
	diff2 := y1 - y2
	return math.Sqrt(math.Pow(diff1, 2) +
		math.Pow(diff2, 2))
}
func main() {}

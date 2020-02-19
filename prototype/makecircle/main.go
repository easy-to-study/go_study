package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	aOpt = flag.Int("a", 0, "Determine the start position of the horizontal axis.")
	bOpt = flag.Int("b", 0, "Determine the start position of the vertical axis.")
)

func main() {
	flag.Parse()

	var r float64
	r = 1000

	var x, y float64

	for i := 0; i < 360; i++ {
		x = float64(*aOpt) + r*math.Cos(math.Pi/180*float64(i))
		y = float64(*bOpt) + r*math.Sin(math.Pi/180*float64(i))

		fmt.Printf("{\"x\"=%v,\"y\"=%v}", x, y)

	}
}

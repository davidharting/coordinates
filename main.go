package main

import (
	"fmt"

	"github.com/davidharting/coordinates/coordinates"
)

func main() {
	p := coordinates.CartesianPair{X: 3, Y: 21}
	fmt.Println(p.Convert())
}

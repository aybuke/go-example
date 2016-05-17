package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	j := i // j int türünden olur
	var x, y int = 5, 12
	var f float64 = math.Sqrt(float64(x*x + y*y)) /*incompatible type açık dönüşüm gerekli*/
	var z int = int(f)
	fmt.Println(x, y, z, j)
}

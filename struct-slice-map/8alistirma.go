package main

import (
	"fmt"
	"math"
)

func main() {
	hypot := func(x, y float64) float64 { /*fonksiyonlar da değer*/
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}

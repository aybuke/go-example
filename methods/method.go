package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 { /*class yok struct var*/
	/*İşaretçi alıcısı ("pointer receiver") kullanmamızın iki nedeni vardır. Birincisi, her metod çağrısında değer kopyalamasından kaçınmak (eğer değer türü büyük bir "struct" ise daha verimli); ikincisi, metodun işaret alıcısının gösterdiği değeri değiştirebilmesi. */

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{7, 24}
	fmt.Println(v.Abs())
}

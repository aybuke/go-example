package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { /*for parantezleri yok + for ; sum < 1000; */
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 { /*while buymuşdu*/
		sum2 += sum2
	}
	fmt.Println(sum2)
}

/* for {	}  --> sonsuzluk*/

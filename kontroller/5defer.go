package main

import "fmt"

func main() {
	defer fmt.Println("world") /* Ertelenmiş fonksiyon çağrısı */

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

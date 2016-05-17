package main

import "fmt"

/*Stringer kendini bir dizgi ("string") olarak tanımlayan bir türdür. fmt paketi değerleri yazdırmak için bu arayüze bakıyormuş*/
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

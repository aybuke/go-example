package main

import "fmt"

type Vertex struct { /*type "name" struct */
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	/*v.X = 4*/
	fmt.Println(v.X)
	p := &v
	p.X = 1e9 /*10^8*/
	fmt.Println(v)
	fmt.Println(p)
}

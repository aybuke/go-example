package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return /*return x not enough arguments to return */
}

var c, python, java bool
/*var i, j int = 1, 2*/

func main() {
	fmt.Println(split(17))
	/*var c, python, java = true, false, "no!"*/
	var i int
	k := 3 /*kısa atama*/
	fmt.Println(i, c, k, python, java) /*ilk değerler 0 false*/
}

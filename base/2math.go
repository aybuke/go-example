package main

import "fmt"
import "math/rand"
import "math"

/*import (
*	"fmt"
*	"math/rand"
*	"math")
*/

func main() {
	fmt.Printf("Now you have %g problems. \n", math.Nextafter(2, 3)) /**/
	fmt.Println("My favorite number is", rand.Intn(10)) /* rand.Seed*/
	fmt.Println(math.Pi) /*math.pi dışa aktarılmamış*/
}

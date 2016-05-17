/*bool
*
*string
*
*int  int8  int16  int32  int64
*uint uint8 uint16 uint32 uint64 uintptr

*byte // diğer adıyla uint8

*rune // diğer adıyla int32
*     // Unicode karakter kodlarını ifade eder
*
*float32 float64
*complex64 complex128 */

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n" /*Type-value*/
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

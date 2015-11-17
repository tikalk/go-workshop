//Covers go basics
//
// Author: Dmitri Krasnenko

package corz

import (
	"fmt"
	"math/cmplx"
)

var (

	a,b int = 5,6
	x int16 = 8
	d int32 = 8
	e float32 = 9.0
	f float64 = 45.9
	maxi uint64 = 1 << 64 - 1
	camp complex128 = cmplx.Sqrt(-5 + 12i)

	my67 = 67
	weAreCool = true
	h interface{} = "This is a string."
)

const (
	c_0 = 56
	c_1 = "This is C1"

	//Won't compile
	//c_2 = [...]int{5}
)

const (
	SUN=iota
	MON
	TUE
)

func Vars()  {
	//Type is static. Won't compile
	//my67 = "String..."

	//Undeclared variable. Won't compile
	//j = []byte{}

	//Ok
	j := true
	j  = false
	fmt.Println(j)

	//Unused variable. Won't compile.
	//k := "Unused variable."
}

func Casts() {
	s := "6"
	i := 42
	u := uint(i)
	f := float64(i)

	//Won't compile, can't convert string to int
	//d := int(s)
	b := []byte(s)
	fmt.Println(b)

	fmt.Println("==== Cast #1 ====")
	fmt.Println("Catsing: ", i, u, f, s)
}

func Switch(){
	var t interface{} = 'c'
	t = true

	fmt.Println("==== Switch #1 ====")
	switch t {
	case 'a':
		fmt.Println("T is 'c'")
	case 'b','c':
		fmt.Println("T is 'b' or 'c'")
	default:
		fmt.Println("T is UFO")
	}

	fmt.Println("==== Switch #2 ====")
	switch t := t.(type) {
	case bool:
		fmt.Printf("T is a bool: %t\n", t)
	case byte:
		fmt.Printf("T is a byte: %t\n", t)
	default:
		fmt.Printf("T is a %T: %v\n", t,t)
	}
}

func Defaults(isDefault bool) (res string) {
	if !isDefault {
		res = "Not a default"
	}

	return "res"
}

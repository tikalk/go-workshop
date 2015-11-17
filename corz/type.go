// Covers maps, slices, arrays
//
// Author: Dmitri Krasnenko

package corz

import (
	"fmt"
	"time"
)

var (
	//Slice
	q = []int{5,6,7}

    //Array
	m = [3]int{5,6,7}

    //Same as above
	n = [...]int{5,6,7}

	l = map[string]interface{}{
		"name": "Alex", "city": "NY",
		"map": map[string]string{"a":"1"}}
)

func AssignAaaaagh() {
	f := q
	f[0] = 89
	fmt.Println(f,q)

	q = append(q,9)
	fmt.Println(f,q)

	arr := [3]int{1,2,3}

	//Won't compile, Statically typed
	//arr = [...]int{4,65}

	//Won't compile, array isn't slice
	//arr = append(arr, 5)

	fmt.Println(arr)
}


func Loops() {
	arr := []string{"a","b","c"}

	fmt.Println("==== Loop #1 ====")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Next string: %s at %d\n", arr[i],i)
	}

	fmt.Println("==== Loop #2 ====")
	i := 0
	for {
		if i >= len(arr) {
			break
		}

		fmt.Printf("Next string: %s at %d\n", arr[i],i); i++
	}


	fmt.Println("==== Loop #3 ====")
	for _,e := range arr {
		fmt.Printf("Next string: %s at %d\n", e)
	}

	fmt.Println("==== Loop #4 ====")
	Outer: for i,e := range arr {
		for c:=0; c<2 ; c++ {
			if i == 2 {
				continue Outer
			}

			fmt.Printf("Next string: %s at %d\n", e,i)
		}
	}
}

func AssignMaps() {
	p := make(map[string]interface{}, 0)

	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("%d", i)
		val := time.Now()
		p[key] = val

		time.Sleep(30)
	}

	fmt.Println(p)

	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("%d", i)
		p[key] = key

		time.Sleep(30)
	}

	fmt.Println(p)

	//Won't compile. Types are different.
	//s := map[string]string{"3":"1","2":"2","1":"3"}; p = s

	s := map[string]interface{}{"3":"1","2":"2","1":"3"}
	p  = s
	fmt.Println(p)

	for key, val := range p {
		fmt.Println(key, val)
	}

	for key, _ := range p {
		p[key] = "KEY"
		p["0"] = "KEY"
		delete(p, "2")

		fmt.Println(key, p[key])
	}

	//Default value if doesn't exist: bool = false, int = 0, string = ""
	k2, ok := p["2"]
	k1, ex := p["1"]
	fmt.Println(k2, ok, k1, ex)
}


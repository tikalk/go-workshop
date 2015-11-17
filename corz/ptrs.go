// Shows when to use pointers
//
// Author: Dmitri Krasnenko

package corz

import "fmt"

type LargeAndExpesiveToCopyType [20]int64

func (l LargeAndExpesiveToCopyType) LargeStringVal() string {
	//NOTE:
	//1. l is passed by value. Meaning, it's a copy of the original
	//2. Any changes made to l object inside the function WON'T affect the original at all
	return fmt.Sprintf("%v", l)
}

func (l *LargeAndExpesiveToCopyType) LargeStringPtr() string {
	//NOTE:
	//1. l is passed by pointer. Meaning, it's the original
	//2. Any changes made to l object inside the function WILL affect the original
	return fmt.Sprintf("%v", *l)
}

type Sequence []int

func (s Sequence) Len() int {
	//NOTE:
	//1. s is a slice
	//2. You may pass a slice by value as it's a structure with small memory usage overhead
	return len(s)
}
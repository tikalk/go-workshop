//Basic example of how to use go routines and channels
//
// Author: Dmitri Krasnenko
package corz

import (
	"fmt"
	"time"
)

func simpleRoutine() {
	fmt.Println("Simple Routine's here!")
}

func Routine() {
	i := 5

	go simpleRoutine()

	go fmt.Println("Print i: ", i)

	go func(a int) {
		fmt.Println("Print a: ", a)
	}(9)

	go func() {
		fmt.Println("Print same i inside: ", i)
	}()

	//Implement function
	printRangeAsync(20)
}

func printRangeAsync(last int) {
	//Write code:
	//1. that generates numbers from 0 to 'last'
	//2. prints them out using a separate go routine
}

func SampleChannel() {
	bw := byte(3)
	ch := make(chan byte, 100)

	go func(wc chan byte) {
		fmt.Println("Byte to channel: ", bw)
		wc <- bw
	}(ch)

	go func(rc chan byte){
		fmt.Println("Byte from channel: ", <- rc)
	}(ch)

	time.Sleep(3 * time.Second)
}
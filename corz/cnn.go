// Shows how to define and work with Go channels.
// Teaches you some tricks and patterns such as circuit breaker etc.
//
// Author: Dmitri Krasnenko

package corz

import (
	"fmt"
	"time"
	_ "math/rand"
)


func SimpleChannel(){
	//Unbuffered channel, sender blocks until receiver has received the value
	chn := make(chan byte)

	msg := "Hi there!"
	for _, c := range []byte(msg) {
		//There are 2 bugs here
		go func(chn chan byte) { chn <- c
		}(chn)
	}

	for i:=0; i<len(msg);i++{
		fmt.Printf("%s", string(<-chn))
	}
}

func receive(ch chan string, done chan struct{}) {
	fmt.Println("Starting receiver...")

	for {
		select {
		case s,more := <- ch:
			if more {
				fmt.Println(s)
				time.Sleep(3 * time.Second)
			}else {
				done <- struct{}{}
				return
			}
		}
	}
}

func SendReceive() {
	//Unbuffered channel
	dn := make(chan struct{})

	//Buffered channel
	//Allows to sender send up to 20 messages without blocking
	rw := make(chan string, 20)

	go sendits(rw, dn)
	go receive(rw, dn)

	<- dn
	fmt.Println("Sender's done.")

	close(rw);
	fmt.Println("Channel's done.")

	<- dn
	fmt.Println("Receiver's done.")
}

func sendits(ch chan string, dn chan struct{}) {
	fmt.Println("Starting sender...")

	buff := []string{"One", "Two", "Three"}
	for _, i := range buff {
		ch <- i
	}

	dn <- struct{}{}
}

func PersonFinder() {
	name := "Jonh Smith"
	fmt.Print("Sending request ... ")
	fmt.Println(<-find_person_deadline(name))
}

func find_person(name string) chan []Person {
	res := make(chan []Person, 1)

	go func(buff chan []Person) {
		//Prepare results. Do some heavy lifting here
		//...
		time.Sleep(10 * time.Second)

		//Send the results to the channel
		res<-[]Person{Person{id:654321098, name:name}}
		close(res)
	}(res)

	return res
}

func find_person_deadline(name string) (res chan []Person){
	res = make(chan []Person, 1)

	go func(res chan []Person) {
		//Create deadline and run the search
		deadline := time.NewTimer(3 * time.Second)
		underlay := find_person(name)

		//Stop deadline on exit and close res channel
		defer close(res)
		defer deadline.Stop()

		select {
		//Receive the results from Person service
		case u := <-underlay:
			res <- u
		//Timeout
		case <-deadline.C:
			res <- []Person{}
		}
	}(res)

	return
}
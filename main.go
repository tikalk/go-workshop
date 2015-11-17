// The main package for the Go workshop.
//
// Author: Dmitri Krasnenko
package main

import (
	"fmt"
	netz "github.com/dsenim/go-workshop/nets"
	core "github.com/dsenim/go-workshop/corz"
)


func main() {
	core.Flag()

	core.Routine()
	core.SampleChannel()

	core.Infoln("My first info log ....")
	core.Warnln("My first warn log ....")
	core.Errnln("My first errn log ....")

	core.Loops()
	core.Casts()
	core.Switch()
	core.AssignMaps()
	core.AssignAaaaagh()

	var l = core.LargeAndExpesiveToCopyType{6,7,8,2}
	fmt.Println("Large by ptr: ", l.LargeStringPtr())
	fmt.Println("Large by val: ", l.LargeStringVal())
	fmt.Println("Large by val: ", l)

	emp := core.NewEmployee(123456, "Adam Smith", "Junior Partner")
	fmt.Println(emp.Name())

	s := core.EmptySet()
	s.Add("a")
	s.Add(4)
	s.Add(7)

	fmt.Println("Set: ", s, s.Size())

	fmt.Println("Item 8 exist: ", s.Exists(8))
	fmt.Println("Item 4 exist: ", s.Exists(4))

	s.Remove(4)
	fmt.Println("Item 4 exist: ", s.Exists(4))

	var lock = &core.Lock{Func:func(s string)(string){return s + "!!!!"}}
	fmt.Println(lock.Func("Lock"))

	var rwLock core.ReadWriteLock = lock
	rwLock.LockForRead()
	rwLock.LockForWrite()

	res := core.Div(4,2)
	div, _ := core.DivSafe(1,0)

	fmt.Println("Result: ",  res)
	fmt.Println("Result: ",  div)

	fmt.Println("Value: ", core.Defaults(true))
	fmt.Println("Value: ", core.Defaults(false))

	core.SendReceive()
	core.SimpleChannel()
	core.PersonFinder()

	netz.Pass()
}

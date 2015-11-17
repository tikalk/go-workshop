// Shows how to use interfaces
//
// Author: Dmitri Krasnenko

package corz

import (
	"fmt"
)

type ReadLock interface {
	LockForRead()
}

type WriteLock interface {
	LockForWrite()
}

type ReadWriteLock interface {
	ReadLock
	WriteLock
}

type F func (i string) (s string)

type Lock struct {
	Func F
}

func NewLock() *Lock {
	return &Lock{echo}
}

func echo(s string) string {
	return s
}

func (l *Lock)  LockForRead() {
	fmt.Println("Locker locked for read.")
}

func (l *Lock)  LockForWrite() {
	fmt.Println("Locker locked for write.")
}


type Visitor interface {
	Visit(o interface{})
}

type Context struct {
	//Won't compile, Visitor is interface
	//log *Visitor

	log Visitor
}

func (c *Context) Log(){
	c.log.Visit(c)
}
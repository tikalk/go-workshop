// Set data structure implementation
//
// Author: Dmitri Krasnenko

package corz

import (
	"fmt"
	"bytes"
)

type Set map[interface{}]bool

func EmptySet() Set {
	return make(Set)
}

func (s Set) Add(i interface{}) {
	s[i]=true
}

func (s Set) Exists(i interface{}) bool {
	return s[i]
}

func (s Set) Remove(i interface{}) {
	delete(s, i)
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) String() string {
	//The coder went a little wild here...
	var buf bytes.Buffer; for i,_ := range s {
		if buf.Len() == 0 {
			buf.WriteString(fmt.Sprintf("%v", i))
		} else {
			buf.WriteString(fmt.Sprintf(", %v", i))
		}
	}

	return "[" + buf.String() + "]"
}